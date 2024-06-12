#include <iostream>
#include <unordered_map>

using namespace std;

int main() {

    int n, q;
    cin >> n >> q;

    int lastGlobalNotification = -1;
    unordered_map<int, int> userNotifications;

    int counter = 1;
    for (int i = 0; i < q; ++i) {

        int t, id;
        cin >> t >> id;

        if (t == 1) {

            if (id == 0) {
                lastGlobalNotification = counter;
                for (auto& user : userNotifications)
                    user.second = counter;
            } else {
                userNotifications[id] = counter;
            }

            ++counter;

        } else {

            if (userNotifications.find(id) != userNotifications.end()) {
                cout << userNotifications[id] << endl;
                continue;
            }

            if (lastGlobalNotification != -1)
                cout << lastGlobalNotification << endl;
            else
                cout << 0 << endl;

        }

    }
    
    return 0;

}
