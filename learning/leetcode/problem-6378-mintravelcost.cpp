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
    int minimumTotalPrice(int n, vector<vector<int>>& edges, vector<int>& price, vector<vector<int>>& trips) {
        unordered_map<int, vector<int>> umaps; //建图
        for(const auto& e: edges) {
            umaps[e[0]].push_back(e[1]);
            umaps[e[1]].push_back(e[0]);
        }
        
        bool ok = false;
        int flag[55] = {0};
        vector<int> vPath; //记录搜索路径
        function<void(int, int)> dfspath = [&](int k, int end) {
            if(ok) return ;

            flag[k] = 1;
            vPath.push_back(k);
            
            if(k == end) {
                ok = true;
                return ;
            }
            for(const auto& z: umaps[k]) {
                if(!flag[z]) {
                    dfspath(z, end);
                    if(ok) return ;
                    vPath.pop_back();
                }
            }
        };

        vector<int> vNums(n, 0);
        
        for(const auto& t: trips) {
            ok = false;
            memset(flag, 0, sizeof(flag));
            vPath.clear();
            dfspath(t[0], t[1]);

            for(int k: vPath) {
                ++vNums[k];
            }
            
            for(int i = 0; i < vPath.size(); ++i) {
                cout << vPath[i] << "->";
            }
            cout << endl;
        }

        int ans = 1e9+7;
        int innersum = 0;
        int maskprice[55] = {0}; //记录对应索引节点价格是否折扣

        function<void(int)> dfscost = [&](int k) {
            if(innersum >= ans) return ;

            if(k >= n) {
                ans = min(ans, innersum);
                // int sum = 0;
                // for(int i = 0; i < n; ++i) {
                //     sum += vNums[i]*price[i]/(maskprice[i]+1);
                // }
                // if(sum < ans) {
                //     ans = sum;
                // }
                return ;
            }

            bool can = true;
            for(int z: umaps[k]) {
                if(maskprice[z]) {
                    can = false;
                    break;
                }
            }
            if(can) {
                innersum += vNums[k]*price[k]/2;
                maskprice[k] = 1;
                dfscost(k+1);
                maskprice[k] = 0;
                innersum -= vNums[k]*price[k]/2;

                innersum += vNums[k]*price[k];
                dfscost(k+1);
                innersum -= vNums[k]*price[k];
            }else {
                innersum += vNums[k]*price[k];
                dfscost(k+1);
                innersum -= vNums[k]*price[k];
            }
        };
        
        dfscost(0);

        return ans;
    }
};
//

int main(int argc, char *argv[]) {
    cout << "~~~~~~template~~~~~~" << endl;

    Solution s;

    vector<vector<int>> c0_edges = {{0, 1}, {1, 2}, {1, 3}};
    vector<int> c0_price = {2, 2, 10, 6};
    vector<vector<int>> c0_trips = {{0, 3}, {2, 1}, {2, 3}};

    cout << s.minimumTotalPrice(4, c0_edges, c0_price, c0_trips) << endl;

    return 0;
}
