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
      <el-tab-pane label="Notifications">
        <el-row>
          <el-col :span="24">
            <el-button type="primary" @click="createNotificationDialogVisible = true">Create Notification</el-button>
          </el-col>
        </el-row>
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
      "fetchProblems",
      "fetchNotifications",
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
  },
};
</script>
