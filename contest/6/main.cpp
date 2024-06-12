#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main() {

    int numTests;
    cin >> numTests;

    for (int i = 0; i < numTests; ++i) {

        int numOrders;
        cin >> numOrders;

        vector<pair<int, int>> arrival(numOrders);
        for (int j = 0; j < numOrders; ++j) {
            cin >> arrival[j].first;
            arrival[j].second = j;
        }
        
        sort(arrival.begin(), arrival.end());

        int numCars;
        cin >> numCars;

        vector<vector<int>> carCharacteristics(numCars, vector<int>(4));
        for (int j = 0; j < numCars; ++j) {
            cin >> carCharacteristics[j][0] >> carCharacteristics[j][1] >> carCharacteristics[j][2];
            carCharacteristics[j][3] = j;
        }

        sort(carCharacteristics.begin(), carCharacteristics.end(), [](const vector<int>& a, const vector<int>& b) {
            if (a[0] == b[0]) return a[3] < b[3];
            return a[0] < b[0];
        });

        // for (auto& [arr, j] : arrival) cout << "(" << arr << ", " << j << ") ";
        // cout << endl;

        // for (auto& charact : carCharacteristics)
        //     cout << "(" << charact[0] << ", " << charact[1] << ", " << charact[2] << ", " << charact[3] << ") ";
        // cout << endl;

        vector<int> carShedule(numOrders, -1);
        for (auto& [arr, idx] : arrival) {

            for (int j = 0; j < carCharacteristics.size(); ++j) {
                if (carCharacteristics[j][0] <= arr && arr <= carCharacteristics[j][1] && carCharacteristics[j][2] > 0) {

                    // cout << "order " << idx << " go to car " << carCharacteristics[j][3] << endl;

                    carShedule[idx] = carCharacteristics[j][3] + 1;
                    --carCharacteristics[j][2];
                    break;
                }
            }

        }

        for (int car : carShedule) cout << car << ' ';
        cout << endl;

    }

    return 0;
    
}