---
title: "设计推特"
date: 2021-04-19T22:04:56+08:00
weight: 1
tags: [排序, 堆]
---

## [355. 设计推特](https://leetcode-cn.com/problems/design-twitter)
设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，  
能够看见关注人（包括自己）的最近十条推文。  
你的设计需要支持以下的几个功能：  
```
postTweet(userId, tweetId): 创建一条新的推文
getNewsFeed(userId): 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。
推文必须按照时间顺序由最近的开始排序。
follow(followerId, followeeId): 关注一个用户
unfollow(followerId, followeeId): 取消关注一个用户
```
```
示例:
Twitter twitter = new Twitter();
// 用户1发送了一条新推文 (用户id = 1, 推文id = 5).
twitter.postTweet(1, 5);
// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
twitter.getNewsFeed(1);
// 用户1关注了用户2.
twitter.follow(1, 2);
// 用户2发送了一个新推文 (推文id = 6).
twitter.postTweet(2, 6);
// 用户1的获取推文应当返回一个列表，其中包含两个推文，id分别为 -> [6, 5].
// 推文id6应当在推文id5之前，因为它是在5之后发送的.
twitter.getNewsFeed(1);
// 用户1取消关注了用户2.
twitter.unfollow(1, 2);
// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
// 因为用户1已经不再关注用户2.
twitter.getNewsFeed(1);
```
## 分析
面向对象的设计，主要逻辑在User类。
### 借助堆
比较难实现的是getNewsFeed， 可用优先队列（小顶堆）来实现，代码略；  
注意优化，从后向前遍历用户的tweet列表，如果优先队列里已经有10个元素，  
当前tweet的时间不大于堆顶tweet的元素，可以不再遍历该用户的tweet列表 。

- main.go
```go
type Twitter struct {
	Users map[int]*User
	time  uint
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{Users: make(map[int]*User, 0)}
}

/** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.getOrAddUser(userId).PostTweet(tweetId, t.time)
	t.time++
}

/** Retrieve the 10 most recent tweet ids in the user's news feed.
Each item in the news feed must be posted by users who the user followed or by the user herself.
Tweets must be ordered from most recent to least recent.
*/
func (t *Twitter) GetNewsFeed(userId int) []int {
	return t.getOrAddUser(userId).GetNewsFeed(t.Users)
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int) {
	t.getOrAddUser(followerId).Follow(followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	t.getOrAddUser(followerId).Unfollow(followeeId)
}

/** search and return a user, if not present, generate and add one **/
func (t *Twitter) getOrAddUser(userId int) *User {
	user, ok := t.Users[userId]
	if !ok {
		user = NewUser(userId)
		t.Users[userId] = user
	}
	return user
}
```
- tweetheap.go
```go
type Heap []*TweetInfo

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Time < h[j].Time }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*TweetInfo))
}
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NewPriorityQueue() *Heap {
	return &Heap{}
}
```
- tweetinfo.go
```go
type TweetInfo struct {
	Id   int
	Time uint
}

func NewTweetInfo(id int, time uint) *TweetInfo {
	return &TweetInfo{Id: id, Time: time}
}
```

- user.go
```go
type User struct {
	tweets    []*TweetInfo
	followees map[int]struct{}
	Id        int
}

func NewUser(id int) *User {
	return &User{Id: id, followees: map[int]struct{}{}}
}

func (u *User) PostTweet(id int, time uint) {
	u.tweets = append(u.tweets, NewTweetInfo(id, time))
}

func (u *User) Follow(id int) {
	if u.Id == id {
		return
	}
	u.followees[id] = struct{}{}
}

func (u *User) Unfollow(id int) {
	delete(u.followees, id)
}

func (u *User) GetNewsFeed(users map[int]*User) []int {
	pq := NewPriorityQueue()
	pushQueue(pq, u.tweets)
	for id := range u.followees {
		if users[id] == nil || len(users[id].tweets) == 0 {
			continue
		}
		pushQueue(pq, users[id].tweets)
	}
	result := make([]int, pq.Len())
	for i := len(result) - 1; i >= 0; i-- {
		result[i] = heap.Pop(pq).(*TweetInfo).Id
	}
	return result
}

func pushQueue(pq *Heap, infos []*TweetInfo) {
	const maxSize = 10
	for i := len(infos) - 1; i >= 0; i-- {
		info := infos[i]
		if pq.Len() < maxSize {
			heap.Push(pq, info)
		} else if top := heap.Pop(pq).(*TweetInfo); top.Time < info.Time {
			heap.Push(pq, info)
		} else {
			heap.Push(pq, top)
			break
		}
	}
}
```

### 排序
也可以直接对所有tweet数组排序，注意到所有数组本来是已经排序的, 只是遍历下从后向前取前10个。

另外，注意一些常见的异常处理，如一个user follow自己。

- main.go
```go
type Twitter struct {
	Users map[int]*User
	time  uint
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{Users: make(map[int]*User, 0)}
}

/** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.time++
	t.getOrAddUser(userId).PostTweet(tweetId, t.time)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed.
Each item in the news feed must be posted by users who the user followed or by the user herself.
Tweets must be ordered from most recent to least recent.
*/
func (t *Twitter) GetNewsFeed(userId int) []int {
	return t.getOrAddUser(userId).GetNewsFeed(t.Users)
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int) {
	t.getOrAddUser(followerId).Follow(followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	t.getOrAddUser(followerId).Unfollow(followeeId)
}

/** search and return a user, if not present, generate and add one **/
func (t *Twitter) getOrAddUser(userId int) *User {
	user, ok := t.Users[userId]
	if !ok {
		user = NewUser(userId)
		t.Users[userId] = user
	}
	return user
}
```
- tweetinfo.go
```go
type TweetInfo struct {
	Id   int
	Time uint
}

func NewTweetInfo(id int, time uint) *TweetInfo {
	return &TweetInfo{Id: id, Time: time}
}
```
- user.go
```go
type User struct {
	tweets    []*TweetInfo
	followees map[int]struct{}
	Id        int
}

func NewUser(id int) *User {
	return &User{Id: id, followees: map[int]struct{}{}}
}

func (u *User) PostTweet(id int, time uint) {
	u.tweets = append(u.tweets, NewTweetInfo(id, time))
}

func (u *User) Follow(id int) {
	if u.Id == id {
		return
	}
	u.followees[id] = struct{}{}
}

func (u *User) Unfollow(id int) {
	delete(u.followees, id)
}

func (u *User) GetNewsFeed(users map[int]*User) []int {
	var tweetsList [][]*TweetInfo
	total := 0
	if len(u.tweets) > 0 {
		total += len(u.tweets)
		tweetsList = append(tweetsList, u.tweets)
	}

	for id := range u.followees {
		if users[id] == nil || len(users[id].tweets) == 0 {
			continue
		}
		tweetsList = append(tweetsList, users[id].tweets)
		total += len(users[id].tweets)
	}
	return getLatest(tweetsList, total)
}

func getLatest(tweetsList [][]*TweetInfo, total int) []int {
	const maxSize = 10
	n := maxSize
	if n > total {
		n = total
	}
	result := make([]int, n)
	for i := 0; i < n; i++ {
		markedRow := -1
		maxTime := uint(0)
		for row, tweets := range tweetsList {
			if len(tweets) == 0 {
				continue
			}
			last := tweets[len(tweets)-1]
			if last.Time > maxTime {
				maxTime = last.Time
				markedRow = row
			}
		}
		if markedRow == -1 {
			continue
		}
		result[i] = tweetsList[markedRow][len(tweetsList[markedRow])-1].Id
		tweetsList[markedRow] = tweetsList[markedRow][:len(tweetsList[markedRow])-1]
	}
	return result
}
```

### 复杂度对比
getNewsFeed复杂度分析：

> 假设n为用户及其关注用户所有tweet的个数， m为用户关注的用户加自己总共的用户数。
> 
> 利用小顶堆的实现： 时间复杂度O(nlogn), 最坏情况下所有tweet都要插入堆里一次； 空间复杂度O(1)， 堆里最多有10个元素。
> 
> 直接遍历数组实现： 时间复杂度O(10*m), 空间复杂度O(m)，有一个中间数组。