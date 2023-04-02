#include <cstdlib>
#include <cstring>
#include <cctype>
#include <cstdio>

#include <iostream>
#include <vector>
#include <list>
#include <queue>
#include <deque>
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
    int findShortestCycle(int n, vector<vector<int>>& edges) {
        unordered_map<int, vector<int>> umape;
        for(const auto& e: edges) {
            umape[e[0]].push_back(e[1]);
            umape[e[1]].push_back(e[0]);
        }

        auto isDead = [](int& k) { return k&0x2; };
        auto setDead = [](int& k) { k |= 0x2; };
        auto isExtd = [](int& k) { return k&0x1; };
        auto setExtd = [](int& k) { k |= 0x1; };
        auto getDist = [](int& k) { return (k&0xFFFF0000)>>16; };
        auto setDist = [](int& k, int d) { k |= (d<<16); };

        vector<int> mark(n, 0);

        int ans = INT_MAX;
        //枚举所有起点+BFS   
        for(int i = 0; i < n; ++i) {
            queue<int, list<int>> q;

            mark.assign(n, 0);
            q.push(i);
            setExtd(mark[i]);
            while(!q.empty()) {
                int t = q.front();
                q.pop();
                setDead(mark[t]);

                int tlen = getDist(mark[t]);
                for(const auto& k: umape[t]) {
                    if(isDead(mark[k])) continue;

                    if(isExtd(mark[k])) {
                        int circlen = getDist(mark[k]) + tlen + 1;
                        ans = min(ans, circlen);
                        goto Label;
                    }else {
                        q.push(k);
                        setExtd(mark[k]);
                        setDist(mark[k], tlen+1);
                    }
                }
            }
            Label:
            {}
        }

        return ans == INT_MAX ? -1 : ans;
    }
};
//

int main(int argc, char *argv[]) {
    cout << "~~~~~~template~~~~~~" << endl;

    Solution s;

    vector<vector<int>> case0_edges = {{0,1},{1,2},{2,0},{3,4},{4,5},{5,6},{6,3}};
    //cout << s.findShortestCycle(7, case0_edges) << endl;

    vector<vector<int>> case1_edges = {{0, 1}, {0, 2}};
    //cout << s.findShortestCycle(4, case1_edges) << endl;

    vector<vector<int>> case2_edges = {{1,3},{3,5},{5,7},{7,1},{0,2},{2,4},{4,0},{6,0},{6,1}};
    cout << s.findShortestCycle(8, case2_edges) << endl;

    return 0;
}
