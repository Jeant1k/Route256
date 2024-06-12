#include <iostream>
#include <vector>
#include <cmath>
#include <algorithm>

using namespace std;

int main() {

    int numTests;
    cin >> numTests;

    for (int i = 0; i < numTests; ++i) {

        int numCars, carCapacity;
        cin >> numCars >> carCapacity;

        int numBoxes;
        cin >> numBoxes;
        vector<int> boxesWeights(numBoxes);

        for (int i = 0; i < numBoxes; ++i) {
            int weight;
            cin >> weight;
            boxesWeights[i] = pow(2, weight);
        }

        sort(boxesWeights.begin(), boxesWeights.end());

        // for (int weight : boxesWeights) cout << weight << ' ';
        // cout << endl;

        int numTransportations = 0, numLoadedCars = 0, curWeight = 0;
        while (!boxesWeights.empty()) {
            if (curWeight + boxesWeights.back() <= carCapacity) {
                curWeight += boxesWeights.back();
                // cout << "Take weight = " << boxesWeights.back() << " curWeight = " << curWeight << endl;
                boxesWeights.pop_back();

                // cout << "Current boxes: ";
                for (int weight : boxesWeights) cout << weight << ' ';
                cout << endl;

            } else {
                for (int j = boxesWeights.size() - 1; j >= 0; --j) {
                    if (curWeight + boxesWeights[j] <= carCapacity) {
                        curWeight += boxesWeights[j];
                        // cout << "Take weight = " << *(boxesWeights.begin() + j) << " curWeight = " << curWeight << endl;
                        boxesWeights.erase(boxesWeights.begin() + j);
                        // cout << "Current boxes: ";
                        for (int weight : boxesWeights) cout << weight << ' ';
                        cout << endl;
                    }
                }
                ++numLoadedCars;
                // cout << "Car is loaded" << endl;
                if (numLoadedCars >= numCars) {
                    ++numTransportations;
                    // cout << "Cars going" << endl;
                    numLoadedCars = 0;
                }
                curWeight = 0;
            }
        }

        if (curWeight > 0 || numLoadedCars > 0) ++numTransportations;

        cout << numTransportations << endl;

    }

    return 0;
    
}