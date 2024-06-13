package main

import (
	"fmt"
	"sort"
)

type Event struct {
	pos   int
	count int
}

func maxJinghuaPosts(n, m, k int, intervals [][2]int) int {
	events := make([]Event, 0, 2*m)

	for _, interval := range intervals {
		li, ri := interval[0], interval[1]
		events = append(events, Event{pos: li, count: 1})  // 精华区间起点
		events = append(events, Event{pos: ri, count: -1}) // 精华区间终点
	}

	// 排序事件，起点在前，终点在后
	sort.Slice(events, func(i, j int) bool {
		if events[i].pos == events[j].pos {
			return events[i].count > events[j].count
		}
		return events[i].pos < events[j].pos
	})

	// 扫描线计算每个位置的精华帖子数量
	currentJinghua := 0
	posJinghua := make([][2]int, 0, len(events))
	for _, event := range events {
		currentJinghua += event.count
		if len(posJinghua) > 0 && posJinghua[len(posJinghua)-1][0] == event.pos {
			posJinghua[len(posJinghua)-1][1] = currentJinghua
		} else {
			posJinghua = append(posJinghua, [2]int{event.pos, currentJinghua})
		}
	}

	// 使用滑动窗口找到长度为 k 的子区间的最大精华帖子数量
	maxJinghuaInWindow := 0
	j := 0
	currentWindowJinghua := 0
	for i := 0; i < len(posJinghua); i++ {
		windowStart := posJinghua[i][0]
		windowEnd := windowStart + k

		// 调整窗口右边界
		for j < len(posJinghua) && posJinghua[j][0] < windowEnd {
			currentWindowJinghua += posJinghua[j][1]
			j++
		}

		// 更新最大精华帖子数量
		if currentWindowJinghua > maxJinghuaInWindow {
			maxJinghuaInWindow = currentWindowJinghua
		}

		// 减去窗口左边界的影响
		currentWindowJinghua -= posJinghua[i][1]
	}

	return maxJinghuaInWindow
}

func main() {
	n := 5
	m := 2
	k := 3
	intervals := [][2]int{
		{1, 2},
		{3, 5},
	}
	fmt.Println(maxJinghuaPosts(n, m, k, intervals)) // 输出结果
}
