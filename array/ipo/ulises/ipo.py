import queue
class Project:
    def __init__(self, value):
        self.value = value
    
    def __lt__(self, next):
        return self.value[0] > next.value[0]
class Solution:
    def findMaximizedCapital(self, k: int, w: int, profits: List[int], capital: List[int]) -> int:
        projects = []
        for i in range(len(profits)):
            p = (profits[i], capital[i])
            projects.append(p)
        projects.sort(key=lambda p: p[1])
        listIndex = 0
        pq = queue.PriorityQueue()
        for i in range(k):
            while listIndex < len(profits) and projects[listIndex][1] <= w:
                pq.put(Project(projects[listIndex]))
                listIndex += 1
            if  not pq.empty():
                w += pq.get().value[0]
        return w