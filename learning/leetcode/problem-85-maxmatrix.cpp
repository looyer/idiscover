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
//case 1. 暴力法，枚举边界
//case 2. 单调栈
class Solution {
public:
    int maximalRectangle(vector<vector<char>>& matrix) {
        if(!matrix.size()) return 0;

        int ncol = matrix[0].size();
        vector<int> vhigh(ncol, 0);

        int ans = 0;
        for(int i = 0; i < matrix.size(); ++i) {
            for(int j = 0; j < ncol; ++j) {
                int c = matrix[i][j] - '0';
                vhigh[j] = (++vhigh[j]) * c;
            }
            for(int j = 0; j < ncol; ++j) {
                int lk = j, rk = j;
                while(lk >= 0 && vhigh[lk] >= vhigh[j]) --lk;
                while(rk < ncol && vhigh[rk] >= vhigh[j]) ++ rk;
                int mh = (rk-lk-1)*vhigh[j];
                ans = max(ans, mh);
            } 
        }

        return ans;
    }
};
//

int main(int argc, char *argv[]) {
    cout << "~~~~~~template~~~~~~" << endl;

    Solution s;

    vector<vector<char>> vvmatrix = {
        {'1','0','1','0','0'},
        {'1','0','1','1','1'},
        {'1','1','1','1','1'},
        {'1','0','0','1','0'},
    };

    cout << s.maximalRectangle(vvmatrix) << endl;

    return 0;
}
