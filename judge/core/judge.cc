#include <filesystem>
#include <iostream>
#include <map>
#include <vector>

using namespace std;
namespace fs = std::filesystem;

int main(int argc, string argv[]) {
	if (argc != 5) {
		cout << "Usage: judge <time_limit> <lang>" << endl;
		return 1;
	}
	Judge judge("/code/", argv[1], argv[2]);

	return 0;
}

class Judge {
   public:
	string workdir, time_limit, lang;
	struct result_info {
		int result, total_case, passed_case, score;
		double time, mem;
		vector<string> info;
	} res;

	Judge(string workdir, string time_limit, string lang)
		: workdir(workdir), time_limit(time_limit), lang(lang) {
		res.total_case = res.passed_case = 0;
		for (int i = 0; fs::exists(workdir + to_string(i) + ".in"); i++) {
			string output, error;
			string cmd;
			if (lang == "c" || lang == "cpp" || lang == "go") {
				cmd = "./main";
			} else if (lang == "java") {
				cmd = "java main";
			} else if (lang == "python") {
				cmd = "python3 main.py";
			}
			run(cmd, workdir + itoa(i) + ".in", output, error);
			if (output.empty()) {
				res.info.push_back("Info: " + to_string(i) + " " + "0.1s" +
								   " " + "0.1MB" + " " + "passed");
				res.passed_case++;
			}
			res.total_case++;
		}
		res.result = res.passed_case == res.total_case ? 1 : 0;
		res.score = 100 * res.passed_case / res.total_case;
	}

   private:
	void run(string cmd, string input, string& output, string& error) {
		cmd = "timeout " + time_limit + " " + cmd + " < " + input +
			  " | diff - " + output;
		string command_output = "";
		string command_error = "";
		FILE* pipe = popen(cmd.c_str(), "r");
		if (!pipe)
			throw runtime_error("popen() failed!");
		char buffer[128];
		while (fgets(buffer, sizeof(buffer), pipe) != NULL) {
			command_output += buffer;
		}
		pclose(pipe);
		output = command_output;
	}

	void init_workdir() {
		system("cp /code/* /main/");
		if (lang == "c")
			system("gcc /main/main.c -o /main/main");
		else if (lang == "cpp")
			system("g++ /main/main.cpp -o /main/main");
		else if (lang == "java")
			system("javac /main/main.java");
		else if (lang == "go")
			system("go build -o /main/main /main/main.go");
	}

	void print_report() {
		/*
		report format:

		Result: 1
		Score: 100
		Passed: 3/4
		Info: case1 0.1s 0.1MB passed
		Info: case2 0.1s 0.1MB passed
		Info: case3 0.1s 0.1MB passed
		Info: case4 0.1s 0.1MB failed
		*/
		cout << "Result: " << res.result << endl
			 << "Score: " << res.score << endl
			 << "Passed: " << res.passed_case << "/" << res.total_case << endl;
		for (auto i : res.info)
			cout << i << endl;
	}
};