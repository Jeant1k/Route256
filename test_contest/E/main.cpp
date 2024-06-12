#include <iostream>
#include <vector>
#include <string>
#include "nlohmann/json.hpp"

using json = nlohmann::json;
using namespace std;

class FileSystemTree {
private:
    class FileSystemNode {
    public:
        string name;
        bool isFile;
        bool isHack;
        vector<FileSystemNode> children;

        FileSystemNode(const string& name, bool isFile = false, bool isHack = false)
            : name(name)
            , isFile(isFile)
            , isHack(isHack)
        {
            if (isFile && name.size() >= 5 && name.substr(name.size() - 5) == ".hack")
                this->isHack = true;
        }
    };

    void printTree(const FileSystemNode& node, int depth = 0) const {

        for (int i = 0; i < depth; ++i) cout << "\t";

        cout << node.name;
        if (node.isFile) cout << " (f)";
        if (node.isHack) cout << " (h)";
        cout << endl;

        for (const auto& child : node.children)
            printTree(child, depth + 1);

    }

    void parseJSON(const json& data, FileSystemNode& node, bool isParentHacked = false) {

        bool isCurDirHacked = isParentHacked;

        if (data.find("files") != data.end()) {
            for (const auto& file : data["files"]) {

                FileSystemNode newFile(file, true, isCurDirHacked);
                isCurDirHacked = newFile.isHack || isCurDirHacked;

                node.children.push_back(newFile);

            }
        }

        for (auto& child : node.children) child.isHack = isCurDirHacked;
        
        if (data.find("folders") != data.end()) {
            for (const auto& folder : data["folders"]) {

                FileSystemNode newFolder(folder["dir"], false, isCurDirHacked);

                parseJSON(folder, newFolder, isCurDirHacked);

                node.children.push_back(newFolder);

            }
        }

    }

    int countHackedFiles(const FileSystemNode& node) const {

        int count = node.isHack && node.isFile ? 1 : 0;

        for (const auto& child : node.children)
            count += countHackedFiles(child);

        return count;

    }

    FileSystemNode root;

public:
    FileSystemTree() : root("root") {}

    void buildTreeFromJSON(int numLines) {

        string line, jsonData = "";
        getline(cin, line);
        for (int j = 0; j < numLines; ++j) {
            getline(cin, line);
            jsonData += line + "\n";
        }
            
        json data = json::parse(jsonData);

        root.name = data["dir"];
        parseJSON(data, root);

    }

    void printTree() const { printTree(root); }

    int countHackedFiles() const { return countHackedFiles(root); }

};

int main() {

    int numTests;
    cin >> numTests;

    for (int i = 0; i < numTests; ++i) {

        int numLines;
        cin >> numLines;

        FileSystemTree tree;

        tree.buildTreeFromJSON(numLines);

        cout << tree.countHackedFiles() << endl;

    }

    return 0;

}
