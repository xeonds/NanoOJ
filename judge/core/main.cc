#include <sys/resource.h>
#include <bits/stdc++.h>
#include <sys/types.h>
#include <sys/wait.h>
#include "sqlite3.h"
#include <unistd.h>
#include <fstream>
#include <vector>

using namespace std;

class SourceCodeJudgement {
public:
    static const int time_limit = 5;
    static bool judge(string source_code, string param, string expect_output) {
        // Compile the source code
        string temp_file = "temp.cpp";
        ofstream temp(temp_file);
        temp << source_code;
        temp.close();
        string compile_command = "g++ -o program " + temp_file;
        int compile_result = system(compile_command.c_str());
        if (compile_result != 0) {
            // Compilation failed
            return false;
        }

        // Run the program with the given parameter
        string run_command = "./program " + param;
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
    
    static void cleanup() {
        system("rm -f program temp.cpp");
    }

    static int run(vector<string> source_files, vector<string> expect_files){
        for (int i = 0; i < source_files.size(); i++) {
            string source_code;
            string expect_output;
            ifstream source_file(source_files[i]);
            ifstream expect_file(expect_files[i]);
            if (source_file.is_open() && expect_file.is_open()) {
                source_code = string((istreambuf_iterator<char>(source_file)), istreambuf_iterator<char>());
                expect_output = string((istreambuf_iterator<char>(expect_file)), istreambuf_iterator<char>());
                source_file.close();
                expect_file.close();
            } else {
                cout << "Error: failed to open source or expect file" << endl;
                return 1;
            }
            if (!SourceCodeJudgement::judge(source_code, "", expect_output)) {
                cout << "Test case " << i+1 << " failed" << endl;
            } else {
                cout << "Test case " << i+1 << " passed" << endl;
            }
        }
        SourceCodeJudgement::cleanup();
        return 0;
    }
};

class NanoOJDB {
public:
    sqlite3 *db;
    char *zErrMsg = 0;
    int rc;
    const char *sql;
    const char* data = "Callback function called";
    static int callback(void *data, int argc, char **argv, char **azColName){
        int i;
        fprintf(stderr, "%s: ", (const char*)data);
        for(i = 0; i<argc; i++){
            printf("%s = %s\n", azColName[i], argv[i] ? argv[i] : "NULL");
        }
        printf("\n");
        return 0;
    }
    NanoOJDB() {
        rc = sqlite3_open("NanoOJ.db", &db);
        if( rc ) {
            fprintf(stderr, "Can't open database: %s\n", sqlite3_errmsg(db));
            exit(0);
        } else {
            fprintf(stderr, "Opened database successfully\n");
        }
    }
    ~NanoOJDB() {
        sqlite3_close(db);
    }
    void execute(const char* sql) {
        rc = sqlite3_exec(db, sql, callback, (void*)data, &zErrMsg);
        if( rc != SQLITE_OK ){
            fprintf(stderr, "SQL error: %s\n", zErrMsg);
            sqlite3_free(zErrMsg);
        } else {
            fprintf(stdout, "Operation done successfully\n");
        }
    }
    void commits_judger() {
        NanoOJDB db;
        const char* sql = "SELECT * FROM commits";
        sqlite3_stmt *stmt;
        sqlite3_prepare_v2(db.db, sql, -1, &stmt, NULL);
        int rc = sqlite3_step(stmt);
        while (rc != SQLITE_DONE) {
            if (rc == SQLITE_ROW) {
                int id = sqlite3_column_int(stmt, 0);
                string source_code = string(reinterpret_cast<const char*>(sqlite3_column_text(stmt, 1)));
                string expect_output = string(reinterpret_cast<const char*>(sqlite3_column_text(stmt, 2)));
                bool result = SourceCodeJudgement::judge(source_code, "", expect_output);
                const char* update_sql = ("UPDATE commits SET result=" + to_string(result) + " WHERE id=" + to_string(id)).c_str();
                db.execute(update_sql);
            }
            rc = sqlite3_step(stmt);
        }
        sqlite3_finalize(stmt);
    }
};


int main(int argc, char * argv[]) {
    bool is_source = true;
    if (argc == 1) {
        cout << "Usage: Single file mode:  ./judge -s [source_file1] [source_file2] ... -e [output_file1] [output_file2] ..." << endl
             << "       SQL query mode:    ./judge -d"<<endl;
        return 1;
    } else if (argc == 2 && string(argv[1]) == "-d") {
        NanoOJDB nano_oj;
    } else {
        vector<string> source_files;
        vector<string> expect_files;
        for (int i = 1; i < argc; i++) {
            string arg = argv[i];
            if (arg == "-s") {
                is_source = true;
            } else if (arg == "-e") {
                is_source = false;
            } else {
                if (is_source) {
                    source_files.push_back(arg);
                } else {
                    expect_files.push_back(arg);
                }
            }
        }
    
        if (source_files.size() != expect_files.size()) {
            cout << "Error: number of source files and expect files do not match" << endl;
            return 1;
        }
    
        return SourceCodeJudgement::run(source_files, expect_files);
    }
    
    return 0;
}
