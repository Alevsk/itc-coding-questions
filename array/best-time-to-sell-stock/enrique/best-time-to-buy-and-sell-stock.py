from sys import maxsize as max_integer_size

class Solution:
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        max_profit = 0
        min_price = max_integer_size
        for i, val in enumerate(prices):
            if val < min_price:
                min_price = val
            elif val - min_price > max_profit:
                max_profit = val - min_price

        return max_profit