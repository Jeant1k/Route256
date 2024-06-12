#include <iostream>
#include <string>

using namespace std;

int main() {

    string sticker;
    cin >> sticker;

    int n;
    cin >> n;

    int start, end;
    string new_sticker;
    for (int i = 0; i < n; ++i) {
        cin >> start >> end >> new_sticker;
        for (int j = start - 1; j < end; ++j)
            sticker[j] = new_sticker[j + 1 - start];
    }

    cout << sticker << endl;
    
    return 0;

}