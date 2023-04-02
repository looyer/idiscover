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
        int ele = 0;
        int num = 0;
    };
    
    long long makeSubKSumEqual(vector<int>& arr, int k) {
        vector<bool> mark(arr.size(), false);
        vector<int>  cpar(arr.size(), 0);
        
        long long ans = 0;
        for(int i = 0; i < arr.size(); ++i) {
            if(mark[i]) continue;
            
            cpar.clear();
            for(int j = i; !mark[j]; j = (j+k)%arr.size()) {
                cpar.push_back(arr[j]);
                mark[j] = true;
            }

            sort(cpar.begin(), cpar.end());
            int m = cpar[cpar.size()/2];
            for(int e: cpar) {
                ans += abs(m-e);
            }
        }
        return ans;
    }
};
//

int main(int argc, char *argv[]) {
    cout << "~~~~~~template~~~~~~" << endl;

    Solution s;
    
    vector<int> case1_v = {1, 4, 1, 3};
    cout << s.makeSubKSumEqual(case1_v, 2) << endl;

    return 0;
}
