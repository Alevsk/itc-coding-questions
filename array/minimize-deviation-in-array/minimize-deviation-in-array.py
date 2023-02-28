class Solution:
    def minimumDeviation(self, nums: List[int]) -> int:
        nodes = []
        for i in range(len(nums)):
            if nums[i] % 2 == 0:
                mi = nums[i]
                while mi % 2 == 0:
                    mi = int(mi / 2)
            else:
                mi = nums[i]
            nodes.append((mi, i))

        sol = max(nodes)[0]- min(nodes)[0]
        heapq.heapify(nodes)
        max_value = max(nodes)[0]
        while True:
            v = nodes[0][0]
            i = nodes[0][1]

            if sol > max_value - v:
                sol = max_value - v

            if v == nums[i] and nums[i] % 2 == 0:
                break
            if v > nums[i]:
                break
            if v * 2 > max_value:
                max_value = v * 2
            heapq.heapreplace(nodes, (v*2, i))
        return sol