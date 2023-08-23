<template>
  <div class="admin-view">
    <el-affix>
      <el-page-header @back="goBack">
        <template #content>
          <el-text type="primary" id="title">X-OJ</el-text><el-text id="subtitle"> | Admin Center</el-text>
        </template>
      </el-page-header>
    </el-affix>
    <el-tabs tab-position="left">
      <el-tab-pane label="Problems">
        <h2>Problems</h2>
        <el-card>
          <template #header>
            <div class="card-header">
              <span>All Problems</span>
              <el-button @click="createProblemDialogVisible = true" type="primary">Create Problem</el-button>
            </div>
          </template>
          <el-table :data="problems" style="width: 100%">
            <el-table-column prop="id" label="ID"></el-table-column>
            <el-table-column prop="title" label="Title"></el-table-column>
            <el-table-column label="Difficulty">
              <template #default="{ row }">
                <el-rate v-model="row.difficulty" :colors="colors" disabled />
              </template>
            </el-table-column>
            <el-table-column prop="CreatedAt" label="Created At"></el-table-column>
            <el-table-column label="Actions">
              <template #default="{ row }">
                <el-button @click="editProblem(row)">Edit</el-button>
                <el-button @click="DeleteProblem(row.id)">Delete</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
        <el-dialog v-model="createProblemDialogVisible" title="Create Problem">
          <el-form :model="newProblem" label-width="100px">
            <el-form-item label="Title">
              <el-input v-model="newProblem.title"></el-input>
            </el-form-item>
            <el-form-item label="Description">
              <el-input v-model="newProblem.description" type="textarea"></el-input>
            </el-form-item>
            <el-form-item label="Inputs">
              <el-input v-model="newProblem.inputs" type="textarea" />
            </el-form-item>
            <el-form-item label="Outputs">
              <el-input v-model="newProblem.outputs" type="textarea" />
            </el-form-item>
            <el-form-item label="Difficulty">
              <el-rate v-model="newProblem.difficulty" :colors="colors" />
            </el-form-item>
          </el-form>
          <div class="dialog-footer">
            <el-button @click="createProblemDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="CreateProblem">Create</el-button>
          </div>
        </el-dialog>
        <el-dialog v-model="editProblemDialogVisible" title="Edit Problem">
          <el-form :model="selectedProblem" label-width="100px">
            <el-form-item label="Title">
              <el-input v-model="selectedProblem.title"></el-input>
            </el-form-item>
            <el-form-item label="Description">
              <el-input v-model="selectedProblem.description" type="textarea"></el-input>
            </el-form-item>
            <el-form-item label="Inputs">
              <el-input v-model="selectedProblem.inputs" type="textarea" />
            </el-form-item>
            <el-form-item label="Outputs">
              <el-input v-model="selectedProblem.outputs" type="textarea" />
            </el-form-item>
            <el-form-item label="Difficulty">
              <el-rate v-model="selectedProblem.difficulty" :colors="colors" />
            </el-form-item>
          </el-form>
          <div class="dialog-footer">
            <el-button @click="editProblemDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="UpdateProblem">Update</el-button>
          </div>
        </el-dialog>
      </el-tab-pane>
      <el-tab-pane label="Contests">
        <h2>Contests</h2>
        <el-card>
          <template #header>
            <div class="card-header">
              <span>All Contests</span>
              <el-button @click="createContestDialogVisible = true" type="primary">Create Contest</el-button>
            </div>
          </template>
          <el-table :data="contests" style="width: 100%">
            <el-table-column prop="id" label="ID"></el-table-column>
            <el-table-column prop="title" label="Title"></el-table-column>
            <el-table-column prop="CreatedAt" label="Created At"></el-table-column>
            <el-table-column label="Actions">
              <template #default="{ row }">
                <el-button @click="editContest(row)">Edit</el-button>
                <el-button @click="DeleteContest(row.id)">Delete</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
        <el-dialog v-model="createContestDialogVisible" title="Create Contest">
          <el-form :model="newContest" label-width="128px">
            <el-form-item label="Title">
              <el-input v-model="newContest.title"></el-input>
            </el-form-item>
            <el-form-item label="Description">
              <el-input v-model="newContest.description" type="textarea"></el-input>
            </el-form-item>
            <el-form-item label="Problems">
              <el-select v-model="newContest.problems" multiple placeholder="Please select">
                <el-option v-for="problem in problems" :key="problem.id" :label="problem.title" :value="problem.id">
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="Time Duration">
              <el-date-picker v-model="newContest.time" type="datetimerange" range-separator="To"
                start-placeholder="Start date" end-placeholder="End date" />
            </el-form-item>
          </el-form>
          <div class="dialog-footer">
            <el-button @click="createContestDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="CreateContest">Create</el-button>
          </div>
        </el-dialog>
        <el-dialog v-model="editContestDialogVisible" title="Edit Contest">
          <el-form :model="selectedContest" label-width="128px">
            <el-form-item label="Title">
              <el-input v-model="selectedContest.title"></el-input>
            </el-form-item>
            <el-form-item label="Description">
              <el-input v-model="selectedContest.description" type="textarea"></el-input>
            </el-form-item>
            <el-form-item label="Problems">
              <el-select v-model="selectedContest.problems" multiple placeholder="Please select">
                <el-option v-for="problem in problems" :key="problem.id" :label="problem.title" :value="problem.id">
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="Time Duration">
              <el-date-picker v-model="selectedContest.time" type="datetimerange" range-separator="To"
                start-placeholder="Start date" end-placeholder="End date" />
            </el-form-item>
          </el-form>
          <div class="dialog-footer">
            <el-button @click="editContestDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="updateContest">Update</el-button>
          </div>
        </el-dialog>
      </el-tab-pane>
      <el-tab-pane label="Notifications">
        <h2>Notifications</h2>
        <el-row>
          <el-col :span="24">
          </el-col>
        </el-row>
        <el-card>
          <template #header>
            <div class="card-header">
              <span>All Notifications</span>
              <el-button type="primary" @click="createNotificationDialogVisible = true">Create
                Notification</el-button>
            </div>
          </template>
          <el-table :data="notifications" style="width: 100%">
            <el-table-column prop="id" label="ID"></el-table-column>
            <el-table-column prop="title" label="通知标题"></el-table-column>
            <el-table-column prop="author" label="发布者"></el-table-column>
            <el-table-column prop="CreatedAt" label="发布时间"></el-table-column>
            <el-table-column label="Actions">
              <template #default="{ row }">
                <el-button @click="editNotification(row)">Edit</el-button>
                <el-button @click="DeleteNotification(row.id)">Delete</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
        <el-dialog v-model="createNotificationDialogVisible" title="Create Notification">
          <el-form :model="newNotification" label-width="100px">
            <el-form-item label="Title">
              <el-input v-model="newNotification.title"></el-input>
            </el-form-item>
            <el-form-item label="Content">
              <el-input v-model="newNotification.content" type="textarea"></el-input>
            </el-form-item>
          </el-form>
          <div class="dialog-footer">
            <el-button @click="createNotificationDialogVisible = false">Cancel</el-button>
            <el-button type="primary"
              @click="createNotification(newNotification); createNotificationDialogVisible = false">Create</el-button>
          </div>
        </el-dialog>
        <el-dialog v-model="editNotificationDialogVisible" title="Edit Notification">
          <el-form :model="selectedNotification" label-width="100px">
            <el-form-item label="Title">
              <el-input v-model="selectedNotification.title"></el-input>
            </el-form-item>
            <el-form-item label="Content">
              <el-input v-model="selectedNotification.content" type="textarea"></el-input>
            </el-form-item>
          </el-form>
          <div class="dialog-footer">
            <el-button @click="editNotificationDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="UpdateNotification">Create</el-button>
          </div>
        </el-dialog>
      </el-tab-pane>
      <el-tab-pane label="Submissions">
        <h2>Submissions</h2>
        <el-card>
          <template #header>
            <div class="card-header">
              <span>All Submissions</span>
            </div>
          </template>
          <el-table :data="submissions" style="width: 100%">
            <el-table-column prop="id" label="ID"></el-table-column>
            <el-table-column prop="problem_id" label="Problem ID"></el-table-column>
            <el-table-column prop="user_id" label="User ID"></el-table-column>
            <el-table-column prop="language" label="Language"></el-table-column>
            <el-table-column prop="status" label="Status"></el-table-column>
            <el-table-column prop="CreatedAt" label="CreatedAt"></el-table-column>
            <el-table-column label="Actions">
              <template #default="{ row }">
                <el-button @click="DeleteSubmission(row.id)">Delete</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-tab-pane>
      <el-tab-pane label="Users">
        <h2>Users</h2>
        <el-table :data="users" style="width: 100%">
          <el-table-column prop="id" label="ID"></el-table-column>
          <el-table-column prop="username" label="Username"></el-table-column>
          <el-table-column prop="email" label="Email"></el-table-column>
          <el-table-column prop="CreatedAt" label="CreatedAt"></el-table-column>
          <el-table-column label="Actions">
            <template #default="{ row }">
              <el-button @click="DeleteUser(row.id)">Delete</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
      <el-tab-pane label="System">
        <h2>System</h2>
        <h3>Danger Zone</h3>
        <el-row>
          <el-col :span="24">
            <el-button type="danger" @click="ClearCache">Clear Cache</el-button>
            <el-button type="danger" @click="Shutdown">Shutdown</el-button>
            <el-button type="danger" @click="Restart">Restart</el-button>
          </el-col>
        </el-row>
      </el-tab-pane>
      <el-tab-pane label="Manual">
        <h2>Manual</h2>
        <p>问题创建：用例之间使用---分割，且输入用例和输出用例数量务必保持一致。</p>
        <p>通知&问题/测试描述：都支持Markdown语法。</p>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";

export default {
  data() {
    return {
      colors: { 2: '#01D842', 4: '#66CCFF', 5: '#FF4040' },
      createProblemDialogVisible: false,
      editProblemDialogVisible: false,
      createNotificationDialogVisible: false,
      editNotificationDialogVisible: false,
      createContestDialogVisible: false,
      editContestDialogVisible: false,
      newProblem: {
        title: "",
        description: "",
        difficulty: 0,
        inputs: '',
        outputs: ''
      },
      selectedProblem: {
        id: "",
        title: "",
        description: "",
        difficulty: 0,
        inputs: '',
        outputs: ''
      },
      newNotification: {
        author: "xeonds",
        title: "",
        content: "",
      },
      selectedNotification: {
        id: "",
        title: "",
        content: "",
      },
      newContest: {
        title: "",
        description: "",
        start_time: "",
        end_time: "",
        time: "",
        problems: [],
      },
      selectedContest: {
        id: "",
        title: "",
        description: "",
        start_time: "",
        end_time: "",
        problems: [],
      },
    };
  },
  computed: {
    ...mapGetters({
      problems: "getProblems",
      notifications: "getNotifications",
      submissions: "getSubmissions",
      _users: "getUsers",
      contests: "getContests",
    }),
    users() {
      //TODO: backend modified, just delete this, rename _users to users
      return this._users.users;
    },
  },
  methods: {
    ...mapActions([
      "createProblem",
      "updateProblem",
      "deleteProblem",
      "fetchProblems",
      "createNotification",
      "updateNotification",
      "deleteNotification",
      "fetchNotifications",
      "createContest",
      "updateContest",
    ]),
    editProblem(problem) {
      this.selectedProblem = { ...problem };
      this.selectedProblem.inputs = this.selectedProblem.inputs.join('\n---\n');
      this.selectedProblem.outputs = this.selectedProblem.outputs.join('\n---\n');
      this.editProblemDialogVisible = true;
    },
    CreateProblem() {
      this.newProblem.inputs = this.newProblem.inputs.split('\n---\n');
      this.newProblem.outputs = this.newProblem.outputs.split('\n---\n');
      this.createProblem(this.newProblem);
      this.createProblemDialogVisible = false;
    },
    UpdateProblem() {
      this.selectedProblem.inputs = this.selectedProblem.inputs.split("\n---\n");
      this.selectedProblem.outputs = this.selectedProblem.outputs.split("\n---\n");
      this.updateProblem(this.selectedProblem);
      this.editProblemDialogVisible = false;
    },
    editNotification(notification) {
      this.selectedNotification = { ...notification };
      this.editNotificationDialogVisible = true;
    },
    UpdateNotification() {
      this.updateNotification(this.selectedNotification);
      this.editNotificationDialogVisible = false;
    },
    CreateContest() {
      this.newContest.start_time = new Date(this.newContest.time[0]).toISOString();
      this.newContest.end_time = new Date(this.newContest.time[1]).toISOString();
      this.newContest.time = undefined;
      this.createContest(this.newContest);
      this.createContestDialogVisible = false;
    },
    editContest(contest) {
      this.selectedContest = { ...contest };
      this.editContestDialogVisible = true;
    },
    UpdateContest() {
      this.updateContest(this.selectedContest);
      this.editContestDialogVisible = false;
    },
    DeleteProblem(id) {
      this.$confirm("This will permanently delete the problem. Continue?", "Warning", {
        confirmButtonText: "OK",
        cancelButtonText: "Cancel",
        type: "warning",
      })
        .then(() => {
          this.deleteProblem(id);
          this.$message({
            type: "success",
            message: "Delete successfully!",
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "Delete canceled",
          });
        });
    },
    DeleteNotification(id) {
      this.$confirm("This will permanently delete the notification. Continue?", "Warning", {
        confirmButtonText: "OK",
        cancelButtonText: "Cancel",
        type: "warning",
      })
        .then(() => {
          this.deleteNotification(id);
          this.$message({
            type: "success",
            message: "Delete successfully!",
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "Delete canceled",
          });
        });
    },
    DeleteUser(id) {
      this.$confirm("This will permanently delete the user. Continue?", "Warning", {
        confirmButtonText: "OK",
        cancelButtonText: "Cancel",
        type: "warning",
      })
        .then(() => {
          this.$http.delete("/users/" + id).then((res) => {
            this.$message({
              type: res.code == 200 ? "success" : "error",
              message: res.code == 200 ? "Delete successfully!" : "Delete failed!",
            });
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "Delete canceled",
          });
        });
    },
    DeleteSubmission(id) {
      this.$confirm("This will permanently delete the submission. Continue?", "Warning", {
        confirmButtonText: "OK",
        cancelButtonText: "Cancel",
        type: "warning",
      })
        .then(() => {
          this.$http.delete("/api/admin/submission/" + id).then((res) => {
            this.$message({
              type: res.code == 200 ? "success" : "error",
              message: res.code == 200 ? "Delete successfully!" : "Delete failed!",
            });
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "Delete canceled",
          });
        });
    },
    DeleteContest(id) {
      this.$confirm("This will permanently delete the contest. Continue?", "Warning", {
        confirmButtonText: "OK",
        cancelButtonText: "Cancel",
        type: "warning",
      })
        .then(() => {
          this.$http.delete(`/contests/${id}`).then((res) => {
            this.$message({
              type: res.code == 200 ? "success" : "error",
              message: res.code == 200 ? "Delete successfully!" : "Delete failed!",
            });
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "Delete canceled",
          });
        });
    },
    goBack() {
      this.$router.go(-1);
    },
    Restart() {
      this.$confirm("This will restart the server. Continue?", "Warning", {
        confirmButtonText: "OK",
        cancelButtonText: "Cancel",
        type: "warning",
      })
        .then(() => {
          this.$http.post("/api/admin/restart").then((res) => {
            this.$message({
              type: res.code == 200 ? "success" : "error",
              message: res.code == 200 ? "Restart successfully!" : "Restart failed!",
            });
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "Restart canceled",
          });
        });
    },
    Shutdown() {
      this.$confirm("This will shutdown the server. Continue?", "Warning", {
        confirmButtonText: "OK",
        cancelButtonText: "Cancel",
        type: "warning",
      })
        .then(() => {
          this.$http.post("/api/admin/shutdown").then((res) => {
            this.$message({
              type: res.code == 200 ? "success" : "error",
              message: res.code == 200 ? "Shutdown successfully!" : "Shutdown failed!",
            });
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "Shutdown canceled",
          });
        });
    },
    ClearCache() {
      this.$confirm("This will clear the cache. Continue?", "Warning", {
        confirmButtonText: "OK",
        cancelButtonText: "Cancel",
        type: "warning",
      })
        .then(() => {
          this.$http.post("/api/admin/clearcache").then((res) => {
            this.$message({
              type: res.code == 200 ? "success" : "error",
              message: res.code == 200 ? "Clear cache successfully!" : "Clear cache failed!",
            });
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "Clear cache canceled",
          });
        });
    },
  },
};
</script>

<style scoped>
.el-page-header {
  padding: 1rem;
  border-bottom: 1px solid #ebeef5;
  text-align: center;
}

#title {
  font-size: 1.2rem;
  font-weight: 600;
}

#subtitle {
  font-size: 1.2rem;
  font-weight: 100;
}

.el-tabs {
  padding: 2rem;
}

.el-tab-pane {
  padding: 2rem;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>