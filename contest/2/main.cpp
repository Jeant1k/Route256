#include <iostream>
#include <unordered_map>

using namespace std;

int main() {

    int n, t;
    cin >> n >> t;

    char letter;
    unordered_map<char, int> pin;
    for (int i = 0; i < n; ++i) {
        cin >> letter;
        ++pin[letter];
    }

    string s;
    for (int i = 0; i < t; ++i) {

        cin >> s;

        unordered_map<char, int> tmp;
        for (char c : s) ++tmp[c];

        bool flag = true;
        for (auto& [letter, count] : pin) {
            if (tmp.find(letter) == tmp.end() || tmp[letter] != pin[letter]) {
                cout << "NO" << endl;
                flag = false;
                break;
            }
        }

        if (flag) cout << "YES" << endl;

    }

    return 0;
    
}