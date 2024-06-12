#include <iostream>
#include <unordered_map>
#include <climits>
#include <cstdlib>
#include <ctime>

using namespace std;

class Trie {
private:
    class Node {
    public:
        unordered_map<char, Node*> children;
        int numSubstr;
        int lastUpdIdx;

        Node(int idx, int num = 1) : numSubstr(num), lastUpdIdx(idx) {}
        ~Node() { for (auto& child : children) delete child.second; }
    };

    void addBlueWord(const string& suffix, int idx, bool isWholeWord) {

        Node* curNode = root;

        for (char letter : suffix) {

            if (curNode->children.find(letter) == curNode->children.end())
                curNode->children[letter] = new Node(idx);

            curNode = curNode->children[letter];

            if (curNode->lastUpdIdx != idx && curNode->numSubstr != INT_MIN) {
                ++curNode->numSubstr;
                curNode->lastUpdIdx = idx;
            }

        }

        if (isWholeWord) curNode->numSubstr = INT_MIN;

    }

    void addRedWord(const string& suffix, int idx, bool isWholeWord) {

        Node* curNode = root;
        bool wentWholeWord = true;

        for (char letter : suffix) {

            if (curNode->children.find(letter) == curNode->children.end()) {
                wentWholeWord = false;
                break;
            }

            curNode = curNode->children[letter];

            if (curNode->lastUpdIdx != idx && curNode->numSubstr != INT_MIN) {
                --curNode->numSubstr;
                curNode->lastUpdIdx = idx;
            }

        }

        if (isWholeWord && wentWholeWord) curNode->numSubstr = INT_MIN;

    }

    void addWhiteWord(const string& word) {

        Node* curNode = root;
        bool wentWholeWord = true;

        for (char letter : word) {

            if (curNode->children.find(letter) == curNode->children.end()) {
                wentWholeWord = false;
                break;
            }

            curNode = curNode->children[letter];

        }

        if (wentWholeWord) curNode->numSubstr = INT_MIN;

    }

    void addBlackWord(const string& suffix, int idx) {

        Node* curNode = root;

        for (char letter : suffix) {

            if (curNode->children.find(letter) == curNode->children.end()) {
                break;
            }

            curNode = curNode->children[letter];

            if (curNode->lastUpdIdx != idx) {
                curNode->numSubstr = INT_MIN;
                curNode->lastUpdIdx = idx;
            }

        }
    }

    void printTrie(Node* node, string prefix = "", string indent = "") {

        if (node == nullptr) return;

        for (auto& child : node->children) {
            cout << indent << child.first << " (" << child.second->numSubstr << ", " << child.second->lastUpdIdx << ")" << endl;
            printTrie(child.second, prefix + child.first, indent + "\t");
        }

    }

    pair<string, int> findMoveWord(Node* node, const string& path = "") {

        if (node == nullptr) return make_pair("", INT_MIN);

        string word = path;
        int maxNumSubstr = node->numSubstr;

        for (auto& child : node->children) {
            pair<string, int> childResult = findMoveWord(child.second, path + child.first);
            if (childResult.second > maxNumSubstr) {
                word = childResult.first;
                maxNumSubstr = childResult.second;
            }
        }

        return make_pair(word, maxNumSubstr);

    }

    Node* root;

public:
    Trie() { root = new Node(INT_MIN, INT_MIN); }
    ~Trie() { delete root; }

    void print() { printTrie(root); }
    pair<string, int> findMoveWord() { return findMoveWord(root); }

    void addSuffix(const string& suffix, char marker, int idx, bool isWholeWord = false) {

        switch (marker) {
        case 'b':
            addBlueWord(suffix, idx, isWholeWord);
            break;
        case 'r':
            addRedWord(suffix, idx, isWholeWord);
            break;
        case 'w':
            addWhiteWord(suffix);
            break;
        case 'x':
            addBlackWord(suffix, idx);
            break;
        default:
            break;
        }

    }
    
};

string genRandStr(int length) {

    srand(time(0));
    string randomString;

    for (int i = 0; i < length; ++i) {
        char randomChar = 'a' + rand() % 26;
        randomString += randomChar;
    }

    return randomString;

}

int main() {

    int numTests;
    cin >> numTests;

    for (int i = 0; i < numTests; ++i) {

        Trie trie;

        int numWords, numBlueWords, numRedWords, idxBlackWord;
        cin >> numWords >> numBlueWords >> numRedWords >> idxBlackWord;

        string word;
        for (int idx = 1; idx <= numWords; ++idx) {

            cin >> word;

            char wordMarker = '\0';
            if (idx <= numBlueWords)
                wordMarker = 'b';
            else if (numBlueWords < idx && idx <= numBlueWords + numRedWords)
                wordMarker = 'r';
            else if (numBlueWords + numRedWords < idx && idx != idxBlackWord)
                wordMarker = 'w';
            else if (numBlueWords + numRedWords < idx && idx == idxBlackWord)
                wordMarker = 'x';
            else cerr << "wordMarker is undefined" << endl;

            trie.addSuffix(word, wordMarker, idx, true);

            if (wordMarker != 'w')
                for (int i = 1; i < word.size(); ++i)
                    trie.addSuffix(word.substr(i), wordMarker, idx);
            
        }

        pair<string, int> moveWord = trie.findMoveWord();

        if (moveWord.second < 0)
            cout << genRandStr(10) << " 0" << endl;
        else
            cout << moveWord.first << ' ' << moveWord.second << endl;

    }

    return 0;

}
