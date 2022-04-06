---
title: "355. 设计推特"
date: 2021-04-19T22:04:56+08:00
weight: 1
tags: [排序, 堆]
---

## [355. 设计推特](https://leetcode-cn.com/problems/design-twitter)

设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，  
能够看见关注人（包括自己）的最近十条推文。  
你的设计需要支持以下的几个功能：  

- `postTweet(userId, tweetId)`
  
  > 创建一条新的推文

- `getNewsFeed(userId)`
  
  > 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。

- `follow(followerId, followeeId)`
  
  > 关注一个用户

- `unfollow(followerId, followeeId)`
  
  > 取消关注一个用户

示例:

> `Twitter twitter = new Twitter()`
> 
> // 用户1发送了一条新推文 (用户id = 1, 推文id = 5).
> 
> `twitter.postTweet(1, 5)`
> 
> // 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
> 
> `twitter.getNewsFeed(1)`
> 
> // 用户1关注了用户2.
> 
> `twitter.follow(1, 2)`
> 
> // 用户2发送了一个新推文 (推文id = 6).
> 
> `twitter.postTweet(2, 6)`
> 
> // 用户1的获取推文应当返回一个列表，其中包含两个推文，id分别为 -> [6, 5].
> 
> // 推文id6应当在推文id5之前，因为它是在5之后发送的.
> 
> `twitter.getNewsFeed(1)`
> 
> // 用户1取消关注了用户2.
> 
> `twitter.unfollow(1, 2)`
> 
> // 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
> 
> // 因为用户1已经不再关注用户2.
> 
> `twitter.getNewsFeed(1)`

## 分析

面向对象的设计，主要逻辑在User类。

> 测试用例有几个很特别，需要注意：
> 
> - 某个用户可能关注/取关自己
> 
> - 每个Api涉及到的用户，可能当前系统里还不存在，需要检查后创建

- main.go

```go
type Twitter struct {
    // 根据要实现的几个Api，要维护所有用户，且能根据id迅速获得用户，用map
    Users map[int]*User
    // 全局时间，每加一条博客自增一次，用于getNewsFeed中的排序
    Time  int
}

func Constructor() Twitter {
    return Twitter{Users: map[int]*User{}}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
    t.ensureUser(userId).PostTweet(tweetId, t.Time)
    t.Time++
}

func (t *Twitter) GetNewsFeed(userId int) []int {
    return t.ensureUser(userId).GetNewsFeed()
}

func (t *Twitter) Follow(followerId int, followeeId int) {
    t.ensureUser(followerId).Follow(t.ensureUser(followeeId))
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
    t.ensureUser(followerId).Unfollow(followeeId)
}

func (t *Twitter) ensureUser(userId int) *User {
    if user, ok := t.Users[userId]; ok {
        return user
    }
    user := &User{ID: userId, Freinds: map[int]*User{}}
    t.Users[userId] = user
    return user
}
```

- tweetheap.go

```go
const Max = 10

type Q struct {
    s []Article
}

func (q *Q) Len() int           { return len(q.s) }
func (q *Q) Less(i, j int) bool { return q.s[i].Time < q.s[j].Time }
func (q *Q) Swap(i, j int)      { q.s[i], q.s[j] = q.s[j], q.s[i] }
func (q *Q) Push(x interface{}) { q.s = append(q.s, x.(Article)) }
func (q *Q) Pop() interface{} {
    old := q.s
    n := len(old)
    res := old[n-1]
    q.s = old[:n-1]
    return res
}
func (q *Q) push(x Article) {
    if q.Len() < Max {
        heap.Push(q, x)
        return
    }
    if q.s[0].Time < x.Time {
        heap.Pop(q)
        heap.Push(q, x)
    }
}
func (q *Q) pop() Article { return heap.Pop(q).(Article) }
```

- user.go

```go
type User struct {
    ID int
    // 发表的博客
    Articles []Article
    // 好友，即关注的用户
    Freinds  map[int]*User
}

type Article struct {
    ID, Time int
}

func (u *User) PostTweet(id, time int) {
    u.Articles = append(u.Articles, Article{id, time})
}

func (u *User) Follow(user *User) {
    if user == nil || user.ID == u.ID {
        return
    }
    u.Freinds[user.ID] = user
}

func (u *User) Unfollow(id int) {
    delete(u.Freinds, id)
}

func (u *User) GetNewsFeed() []int {
    q := &Q{}
    // 每个人的博客最多取最近发表的10条
    getLastMax := func(articles []Article) []Article {
        from := max(0, len(articles)-Max)
        return articles[from:]
    }
    for _, v := range getLastMax(u.Articles) {
        q.push(v)
    }
    for _, f := range u.Freinds {
        if f == nil {
            continue
        }
        for _, v := range getLastMax(f.Articles) {
            q.push(v)
        }
    }
    res := make([]int, q.Len())
    for i := len(res) - 1; i >= 0; i-- {
        res[i] = q.pop().ID
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

`getNewsFeed` 复杂度分析：

假设n为用户及其关注用户所有tweet的个数， 时间复杂度`O(nlogn)`, 最坏情况下所有tweet都要插入堆里一次； 空间复杂度`O(1)`， 堆里最多有10个元素。