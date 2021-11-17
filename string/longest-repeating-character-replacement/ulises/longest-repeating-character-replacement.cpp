class Solution {
public:
    int characterReplacement(string s, int k) {
        int l = s.length();
        int sol = 0;
        for(int i = 0; i + sol - 1 < l; i++)
        {
            int charCount[30] = { 0 };
            int max = 0;
            for(int j = i; j < l; j++)
            {
                int index = s[j] - 'A';
                charCount[index]++;
                if(charCount[index] > max)
                    max = charCount[index];                    
                int subStringLength = j - i + 1;
                if(subStringLength - max <= k)
                    if(subStringLength > sol)
                        sol = subStringLength;
            }
        }
        return sol;      
    }
};