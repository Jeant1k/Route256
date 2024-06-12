#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

double countStrategy3(vector<vector<double>>& cources, int first, int second, int third) {

    double res = 0;

    for (int i = 0; i < 3; ++i) {
        for (int j = 0; j < 3; ++j) {
            for (int k = 0; k < 3; ++k) {
                if (i != j && j != k && i != k) {
                    double sum = cources[i][first] * cources[j][second] * cources[k][third];
                    cout << i << j << k << ": " << sum << endl;
                    res = max(res, sum);
                }
            }
        }
    }

    cout << "strategy " << first << second << third << ": " << res << endl;

    return res;

}

double countStrategy2(vector<vector<double>>& cources, int first, int second) {

    double res = 0;

    for (int i = 0; i < 3; ++i) {
        for (int j = 0; j < 3; ++j) {
            if (i != j) {
                double sum = cources[i][first] * cources[j][second];
                cout << i << j << ": " << sum << endl;
                res = max(res, sum);
            }
        }
    }

    cout << "strategy " << first << second << ": " << res << endl;

    return res;

}

double countStrategy1(vector<vector<double>>& cources, int first) {

    double res = 0;

    for (int i = 0; i < 3; ++i) {
        double sum = cources[i][first];
        cout << i << ": " << sum << endl;
        res = max(res, sum);
    }

    cout << "strategy " << first << ": " << res << endl;

    return res;

}

int main() {

    int t;
    cin >> t;

    for (int i = 0; i < t; ++i) {

        vector<vector<double>> cources(3, vector<double>(6));
        for (int j = 0; j < 3; ++j) {
            for (int k = 0; k < 6; ++k) {
                double n, m;
                cin >> n >> m;
                cources[j][k] = m / n;
            }
        }

        for (auto& bank : cources) {
            for (auto& cource : bank) {
                cout << cource << ' ';
            }
            cout << endl;
        }

        vector<double> results = {countStrategy3(cources, 0, 2, 0), countStrategy3(cources, 0, 3, 5), countStrategy3(cources, 1, 4, 0), countStrategy2(cources, 1, 5), countStrategy1(cources, 0)};

        cout << *max_element(results.begin(), results.end()) << endl;

    }

    return 0;
    
}