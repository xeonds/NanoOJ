<template>
  <div class="admin-view">
    <el-tabs tab-position="left">
      <el-tab-pane label="Problems">
        <el-row>
          <el-col :span="24">
            <el-button type="primary" @click="createProblemDialogVisible = true"
              >Create Problem</el-button
            >
          </el-col>
        </el-row>
        <el-table :data="problems" style="width: 100%">
          <el-table-column prop="id" label="ID"></el-table-column>
          <el-table-column prop="title" label="Title"></el-table-column>
          <el-table-column
            prop="difficulty"
            label="Difficulty"
          ></el-table-column>
          <el-table-column
            prop="created_at"
            label="Created At"
          ></el-table-column>
          <el-table-column label="Actions">
            <template #default="{ row }">
              <el-button type="text" @click="editProblem(row)">Edit</el-button>
              <el-button type="text" @click="deleteProblem(row)"
                >Delete</el-button
              >
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
            <el-button @click="createProblemDialogVisible = false"
              >Cancel</el-button
            >
            <el-button type="primary" @click="createProblem(newProblem)"
              >Create</el-button
            >
          </div>
        </el-dialog>
        <el-dialog v-model="editProblemDialogVisible" title="Edit Problem">
          <el-form :model="selectedProblem" label-width="100px">
            <el-form-item label="Title">
              <el-input v-model="selectedProblem.title"></el-input>
            </el-form-item>
            <el-form-item label="Difficulty">
              <el-select v-model="selectedProblem.difficulty">
                <el-option label="Easy" value="easy"></el-option>
                <el-option label="Medium" value="medium"></el-option>
                <el-option label="Hard" value="hard"></el-option>
              </el-select>
            </el-form-item>
          </el-form>
          <div class="dialog-footer">
            <el-button @click="editProblemDialogVisible = false"
              >Cancel</el-button
            >
            <el-button type="primary" @click="updateProblem">Update</el-button>
          </div>
        </el-dialog>
      </el-tab-pane>
      <el-tab-pane label="Notifications">
        <el-row>
          <el-col :span="24">
            <el-button
              type="primary"
              @click="createNotificationDialogVisible = true"
              >Create Notification</el-button
            >
          </el-col>
        </el-row>
        <el-table :data="notifications" style="width: 100%">
          <el-table-column prop="id" label="ID"></el-table-column>
          <el-table-column
            prop="problem_title"
            label="通知标题"
          ></el-table-column>
          <el-table-column prop="user_name" label="发布者"></el-table-column>
          <el-table-column prop="status" label="Status"></el-table-column>
          <el-table-column prop="created_at" label="发布时间"></el-table-column>
          <el-table-column label="Actions">
            <template #default="{ row }">
              <el-button type="text" @click="viewSubmission(row)"
                >View</el-button
              >
            </template>
          </el-table-column>
        </el-table>
        <el-dialog
          v-model="createNotificationDialogVisible"
          title="Create Notification"
        >
          <el-form :model="newNotification" label-width="100px">
            <el-form-item label="Title">
              <el-input v-model="newNotification.title"></el-input>
            </el-form-item>
            <el-form-item label="Content">
              <el-input v-model="newNotification.content"></el-input>
            </el-form-item>
            <el-form-item label="Release Time">
              <el-col :span="11">
                <el-date-picker
                  v-model="newNotification.release"
                  type="date"
                  placeholder="Pick a date"
                  style="width: 100%"
                />
              </el-col>
            </el-form-item>
            <el-form-item label="Update Time">
              <el-col :span="11">
                <el-date-picker
                  v-model="newNotification.updateTime"
                  type="date"
                  placeholder="Pick a date"
                  style="width: 100%"
                />
              </el-col>
            </el-form-item>
          </el-form>
          <div class="dialog-footer">
            <el-button @click="createNotificationDialogVisible = false"
              >Cancel</el-button
            >
            <el-button
              type="primary"
              @click="createNotification(newNotification)"
              >Create</el-button
            >
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
                <el-date-picker
                  v-model="newNotification.release"
                  type="date"
                  placeholder="Pick a date"
                  style="width: 100%"
                />
              </el-col>
            </el-form-item>
          </el-form>
          <div class="dialog-footer">
            <el-button @click="createNotificationDialogVisible = false"
              >Cancel</el-button
            >
            <el-button
              type="primary"
              @click="updateNotification(newNotification)"
              >Create</el-button
            >
          </div>
        </el-dialog>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";

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
        difficulty: "",
      },
      newNotification: {
        id: "123",
        author: "xeonds",
        release: "",
        updateTime: "",
        title: "",
        content: "",
      },
    };
  },
  computed: {
    ...mapState({
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
  },
};
</script>
