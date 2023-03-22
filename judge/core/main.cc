#include <sys/resource.h>
#include <bits/stdc++.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <unistd.h>
#include <fstream>
#include <vector>

using namespace std;

class SourceCodeJudgement {
private:
    vector<string> source_files;
    vector<string> input_files;
    vector<string> expect_files;
    int time_limit;

    bool judge(string source_file, string input_file, string expect_file) {
        string compile_command = "g++ -o program " + source_file;
        string expect_output;
        ifstream Expect_file(expect_file);
        if (Expect_file.is_open()) {
            expect_output = string((istreambuf_iterator<char>(Expect_file)), istreambuf_iterator<char>());
        }
        Expect_file.close();

        int compile_result = system(compile_command.c_str());
        if (compile_result != 0) {
            // Compilation failed
            return false;
        }

        // Run the program with the given input
        string run_command = "./program < " + input_file;
        string program_output;
        FILE *fp = popen(run_command.c_str(), "r");
        char buf[1024];
        while (fgets(buf, 1024, fp)) {
            program_output += buf;
        }
        pclose(fp);

        // Compare the program output with the expected output
        program_output.erase(std::remove(program_output.begin(), program_output.end(), '\r'), program_output.end());
        expect_output.erase(std::remove(expect_output.begin(), expect_output.end(), '\r'), expect_output.end());
        program_output.erase(std::remove(program_output.begin(), program_output.end(), '\n'), program_output.end());
        expect_output.erase(std::remove(expect_output.begin(), expect_output.end(), '\n'), expect_output.end());
        return program_output == expect_output;
    }

public:
    SourceCodeJudgement(vector<string> source_files, vector<string> input_files, vector<string> expect_files, int time_limit = 5) : source_files(source_files), input_files(input_files), expect_files(expect_files), time_limit(time_limit) {}

    ~SourceCodeJudgement() {
        remove("program");
        // for (int i = 0; i < source_files.size(); i++) {
        //     remove(source_files[i].c_str());
        //     remove(input_files[i].c_str());
        //     remove(expect_files[i].c_str());
        // }
    }

    int run(){
        for (int i = 0; i < source_files.size(); i++) {
            if (!judge(source_files[i], input_files[i], expect_files[i])) {
                cout << "Test case " << i+1 << " failed" << endl;
            } else {
                cout << "Test case " << i+1 << " passed" << endl;
            }
        }
        return 0;
    }
};

int main(int argc, char * argv[]) {
    if (argc == 1) {
        cout << "Usage: Single file mode:  ./judge -s [source_file1] ... -i [input_file1] ... -e [output_file1] ..." << endl;
        return 1;
    } else {
        // Read input from params
        vector<string> source_files;
        vector<string> input_files;
        vector<string> expect_files;
        int flag = 0;
        for (int i = 1; i < argc; i++) {
            string arg = argv[i];
            if (arg == "-s") {
                flag = 0;
            } else if (arg == "-i") {
                flag = 1;
            } else if (arg == "-e") {
                flag = 2;
            } else {
                if (flag == 0) {
                    source_files.push_back(arg);
                } else if (flag == 1) {
                    input_files.push_back(arg);
                } else{
                    expect_files.push_back(arg);
                }
            }
        }
        if (source_files.size() != expect_files.size() || source_files.size() != input_files.size() || expect_files.size() != input_files.size()) {
            cout << "Error: number of source files, input files and expect files do not match" << endl;
            return 1;
        }
        SourceCodeJudgement judge(source_files, input_files, expect_files);
        judge.run();
    }
    
    return 0;
}
