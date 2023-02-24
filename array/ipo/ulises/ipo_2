import heapq
class Solution:
    def findMaximizedCapital(self, k: int, w: int, profits: List[int], capital: List[int]) -> int:
        projects = []
        max_profit = 10000
        for i in range(len(profits)):
            p = (max_profit - profits[i], capital[i])
            projects.append(p)
        projects.sort(key=lambda p: p[1])
        listIndex = 0
        pq = []
        for i in range(k):
            while listIndex < len(profits) and projects[listIndex][1] <= w:
                heappush(pq, projects[listIndex])
                listIndex += 1
            if len(pq):
                w += max_profit - heappop(pq)[0]
        return w
