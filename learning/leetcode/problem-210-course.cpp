//https://leetcode.cn/problems/course-schedule-ii/

#include <cstdlib>
#include <cstring>
#include <cctype>
#include <cstdio>

#include <iostream>
#include <vector>
#include <map>
#include <set>
#include <unordered_map>
#include <unordered_set>
#include <functional>
#include <algorithm>
#include <memory>
#include <string>

using namespace std;

//
class Solution {
public:
    vector<int> findOrder(int numCourses, vector<vector<int>>& prerequisites) {
        unordered_map<int, vector<int>> umap_succ;
        for(const auto& v: prerequisites) {
            umap_succ[v[1]].push_back(v[0]);
        }

        const static int max_course = 100005;
        int flag[max_course] = {0}; //记未搜索为0，搜索中为-1， 搜索完成为对应栈编号
        int circle = 0; //是否存在环
        int stkid = 1; //入栈编号

        function<void(int)> dfs = [&](int k) {
            if(circle) return ;

            auto iter = umap_succ.find(k);
            if(iter == umap_succ.end()) {
                flag[k] = stkid++;
                return ;
            }

            for(const auto& n: iter->second) {
                if(flag[n] == 0) {
                    flag[n] = -1;
                    dfs(n);
                }else if(flag[n] == -1) {
                    circle = 1; //有环
                    return ;
                }
            }
            flag[k] = stkid++;
        };

        vector<int> ans;

        for(int i = 0; i < numCourses; ++i) {
            if(flag[i] == 0) {
                dfs(i);
            }
            if(circle) return ans;
        }
        ans.resize(numCourses);
        for(int i = 0; i < numCourses; ++i) {
            ans[flag[i]-1] = i;
        }
        std::reverse(ans.begin(), ans.end());
        return ans;
    }
};
//

int main(int argc, char *argv[]) {
    cout << "~~~~~~template~~~~~~" << endl;

    Solution s;

    vector<vector<int>> prep0 = {{1, 0}};
    vector<int> ans = s.findOrder(2, prep0);
    for(const auto& v: ans) {
        cout << v << ", ";
    }
    cout << endl;

    vector<vector<int>> prep1 = {{1, 0}, {2, 0}, {3, 1}, {3, 2}};
    ans = s.findOrder(4, prep1);
        for(const auto& v: ans) {
        cout << v << ", ";
    }
    cout << endl;

    return 0;
}
