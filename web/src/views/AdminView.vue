<template>
  <div class="admin-view">
    <el-tabs tab-position="left">
      <el-tab-pane label="Problems">
        <el-row>
          <el-col :span="24">
            <el-button type="primary" @click="createProblemDialogVisible = true">Create Problem</el-button>
          </el-col>
        </el-row>
        <el-table :data="problems" style="width: 100%">
          <el-table-column prop="ProblemID" label="ID"></el-table-column>
          <el-table-column prop="ProblemTitle" label="Title"></el-table-column>
          <el-table-column prop="ProblemDifficulty" label="Difficulty"></el-table-column>
          <el-table-column prop="CreatedAt" label="Created At"></el-table-column>
          <el-table-column label="Actions">
            <template #default="{ row }">
              <el-button @click="editProblem(row)">Edit</el-button>
              <el-button @click="deleteProblem(row)">Delete</el-button>
            </template>
          </el-table-column>
        </el-table>
        <el-dialog v-model="createProblemDialogVisible" title="Create Problem">
          <el-form :model="newProblem" label-width="100px">
            <el-form-item label="Title">
              <el-input v-model="newProblem.title"></el-input>
            </el-form-item>
            <el-form-item label="Difficulty">
              <el-select v-model="newProblem.difficulty">
                <el-option label="Easy" value="easy"></el-option>
                <el-option label="Medium" value="medium"></el-option>
                <el-option label="Hard" value="hard"></el-option>
              </el-select>
            </el-form-item>
          </el-form>
          <div class="dialog-footer">
            <el-button @click="createProblemDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="createProblem(newProblem)">Create</el-button>
          </div>
        </el-dialog>
        <el-dialog v-model="editProblemDialogVisible" title="Edit Problem">
          <el-form :model="selectedProblem" label-width="100px">
            <el-form-item label="Title">
              <el-input v-model="selectedProblem.ProblemTitle"></el-input>
            </el-form-item>
            <el-form-item label="Description">
              <el-input v-model="selectedProblem.ProblemDescription" type="textarea"></el-input>
            </el-form-item>
            <el-form-item label="Inputs">
              <el-input v-model="selectedProblem.inputs" type="textarea" />
            </el-form-item>
            <el-form-item label="Outputs">
              <el-input v-model="selectedProblem.outputs" type="textarea" />
            </el-form-item>
            <el-form-item label="Difficulty">
              <el-select v-model="newProblem.difficulty">
                <el-option label="Easy" value="easy"></el-option>
                <el-option label="Medium" value="medium"></el-option>
                <el-option label="Hard" value="hard"></el-option>
              </el-select>
            </el-form-item>
          </el-form>
          <div class="dialog-footer">
            <el-button @click="editProblemDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="UpdateProblem">Update</el-button>
          </div>
        </el-dialog>
      </el-tab-pane>
      <el-tab-pane label="Notifications">
        <el-row>
          <el-col :span="24">
            <el-button type="primary" @click="createNotificationDialogVisible = true">Create Notification</el-button>
          </el-col>
        </el-row>
        <el-table :data="notifications" style="width: 100%">
          <el-table-column prop="ID" label="ID"></el-table-column>
          <el-table-column prop="Title" label="通知标题"></el-table-column>
          <el-table-column prop="Author" label="发布者"></el-table-column>
          <el-table-column prop="CreatedAt" label="发布时间"></el-table-column>
          <el-table-column label="Actions">
            <template #default="{ row }">
              <el-button @click="editNotification(row)">Edit</el-button>
              <el-button @click="deleteNotification(row)">Delete</el-button>
            </template>
          </el-table-column>
        </el-table>
        <el-dialog v-model="createNotificationDialogVisible" title="Create Notification">
          <el-form :model="newNotification" label-width="100px">
            <el-form-item label="Title">
              <el-input v-model="newNotification.title"></el-input>
            </el-form-item>
            <el-form-item label="Content">
              <el-input v-model="newNotification.content"></el-input>
            </el-form-item>
            <el-form-item label="Release Time">
              <el-col :span="11">
                <el-date-picker v-model="newNotification.release" type="date" placeholder="Pick a date"
                  style="width: 100%" />
              </el-col>
            </el-form-item>
            <el-form-item label="Update Time">
              <el-col :span="11">
                <el-date-picker v-model="newNotification.updateTime" type="date" placeholder="Pick a date"
                  style="width: 100%" />
              </el-col>
            </el-form-item>
          </el-form>
          <div class="dialog-footer">
            <el-button @click="createNotificationDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="createNotification(newNotification)">Create</el-button>
          </div>
        </el-dialog>
        <el-dialog v-model="editNotificationDialogVisible" title="Edit Problem">
          <el-form :model="newNotification" label-width="100px">
            <el-form-item label="Title">
              <el-input v-model="newNotification.title"></el-input>
            </el-form-item>
            <el-form-item label="Content">
              <el-input v-model="newNotification.content"></el-input>
            </el-form-item>
            <el-form-item label="Release Time">
              <el-col :span="11">
                <el-date-picker v-model="newNotification.release" type="date" placeholder="Pick a date"
                  style="width: 100%" />
              </el-col>
            </el-form-item>
          </el-form>
          <div class="dialog-footer">
            <el-button @click="createNotificationDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="updateNotification(newNotification)">Create</el-button>
          </div>
        </el-dialog>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";

export default {
  data() {
    return {
      createProblemDialogVisible: false,
      editProblemDialogVisible: false,
      createNotificationDialogVisible: false,
      editNotificationDialogVisible: false,
      newProblem: {
        title: "",
        difficulty: "easy",
      },
      selectedProblem: {
        id: "",
        title: "",
        description: "",
        difficulty: "",
        inputs: "",
        outputs: ""
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
    };
  },
  computed: {
    ...mapGetters({
      problems: "getProblems",
      notifications: "getNotifications",
    }),
  },
  methods: {
    ...mapActions([
      "createProblem",
      "updateProblem",
      "deleteProblem",
      "createNotification",
      "updateNotification",
      "deleteNotification",
    ]),
    editProblem(problem) {
      this.selectedProblem = { ...problem };
      this.editProblemDialogVisible = true;
    },
    CreateProblem() {
      this.newProblem.inputs = this.newProblem.inputs.split("\n---\n");
      this.newProblem.outputs = this.newProblem.outputs.split("\n---\n");
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
    deleteProblem(problem) {
      this.$confirm("This will permanently delete the problem. Continue?", "Warning", {
        confirmButtonText: "OK",
        cancelButtonText: "Cancel",
        type: "warning",
      })
        .then(() => {
          this.deleteProblem(problem.ID);
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
  },
};
</script>
