#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main() {

    int n, m;
    cin >> n >> m;

    vector<pair<int, int>> cards(n);
    for (int i = 0; i < n; ++i) {
        cin >> cards[i].first;
        cards[i].second = i;
    }
    
    stable_sort(cards.begin(), cards.end());

    int curCard = 0;
    vector<int> presents(n, -1);
    for (int i = 1; i <= m; ++i) {
        if (curCard > n) break;
        if (i > cards[curCard].first) {
            presents[cards[curCard].second] = i;
            ++curCard;
        }
    }

    for (int present : presents) {
        if (present == -1) {
            cout << -1 << endl;
            return 0;
        }
    }

    for (int present : presents)
        cout << present << ' ';
    cout << endl;
    
    return 0;

}