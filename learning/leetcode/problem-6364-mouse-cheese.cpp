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
    struct node {
        int r = 0; //第一只吃的价值
        int s = 0; //第二只吃的价值
    };
    
    int miceAndCheese(vector<int>& reward1, vector<int>& reward2, int k) {
        vector<node> vn;
        for(int i = 0; i < reward1.size(); ++i) {
            vn.push_back(node{.r = reward1[i], .s = reward2[i], });
        }
        sort(vn.begin(), vn.end(), [](const auto& lhs, const auto& rhs) {
            return (lhs.r + rhs.s > lhs.s + rhs.r) || 
            ((lhs.r + rhs.s == lhs.s + rhs.r) && (lhs.r + lhs.s > rhs.r + rhs.s));
        });
        int ans = 0;
        for(int i = 0; i < vn.size(); ++i) {
            ans += (i < k) ? vn[i].r : vn[i].s;
        }
        return ans;
    }
};
//

int main(int argc, char *argv[]) {
    cout << "~~~~~~template~~~~~~" << endl;

    Solution s;

    vector<int> case0_r1 = {11,14,4,8,7,6,2,7,10,1,2,7,8,4,9,5,13,9,13,9,4,8,10,10,3,1};
    vector<int> case0_r2 = {5,8,4,4,8,6,3,11,5,15,7,2,13,12,8,2,7,3,14,9,5,9,4,6,7,2};
    int case0_k = 5;

    cout << case0_r1.size() << "  " << case0_r2.size() << endl;
    cout << s.miceAndCheese(case0_r1, case0_r2, case0_k) << endl;

    return 0;
}
