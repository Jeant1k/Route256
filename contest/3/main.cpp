#include <iostream>
#include <string>

using namespace std;

int main() {

    int t;
    cin >> t;

    for (int i = 0; i < t; ++i) {

        string s;
        cin >> s;

        for (int i = 1; i < s.size() - 1; ++i) {
            if (s[i - 1] != s[i] && s[i] != s[i + 1])
                s = s.substr(0, i) + s.substr(i + 1);
        }

        bool flag = true;
        for (int i = 0; i < s.size() - 1; ++i) {
            if (s[i] != s[i + 1]) {
                flag = false;
                break;
            }
        }

        if (flag) cout << "YES" << endl;
        else cout << "NO" << endl;
        
    }

    return 0;
    
}