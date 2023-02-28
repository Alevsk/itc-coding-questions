class Project {
    public:
        int profit, capital;
};
bool compareProjectsByCapital(const Project &a, const Project &b)
{
   return a.capital < b.capital;
}
struct CompareProjects {
    bool operator()(const Project& a, const Project& b) {
        return a.profit < b.profit;
    }
};
class Solution {
public:
    int findMaximizedCapital(int k, int w, vector<int>& profits, vector<int>& capital) {
        int size = profits.size();
        vector<Project> projects;
        for(int i = 0; i < size; i++)
        {
            Project p;
            p.profit = profits[i];
            p.capital = capital[i];
            projects.push_back(p);
        }
        sort(projects.begin(), projects.end(), compareProjectsByCapital);
        int totalCapital = w;
        priority_queue<Project, vector<Project>, CompareProjects> pq;
        int vectorIndex = 0;
        for(int i = 0; i < k; i++)
        {
            while(vectorIndex < size && projects[vectorIndex].capital <= totalCapital)
            {
                pq.push(projects[vectorIndex++]);
            }
            if(!pq.empty())
            {
                totalCapital += pq.top().profit;
                pq.pop();
            }
        }
        return totalCapital;
    }
};