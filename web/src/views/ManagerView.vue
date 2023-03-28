<template>
  <div class="manager-view">
    
    <el-tabs v-model="activeTab">
      <el-tab-pane label="Problems" name="problems">
        <el-row>
          <el-col :span="24">
            <el-button type="primary" @click="createProblemDialogVisible = true">Create Problem</el-button>
          </el-col>
        </el-row>
        <el-table :data="problems" style="width: 100%">
          <el-table-column prop="id" label="ID"></el-table-column>
          <el-table-column prop="title" label="Title"></el-table-column>
          <el-table-column prop="difficulty" label="Difficulty"></el-table-column>
          <el-table-column prop="created_at" label="Created At"></el-table-column>
          <el-table-column prop="updated_at" label="Updated At"></el-table-column>
          <el-table-column label="Actions">
            <template #default="{ row }">
              <el-button type="text" @click="editProblem(row)">Edit</el-button>
              <el-button type="text" @click="deleteProblem(row)">Delete</el-button>
            </template>
          </el-table-column>
        </el-table>
        <el-dialog :visible="createProblemDialogVisible" @update:visible="val => createProblemDialogVisible = val" title="Create Problem">
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
          <div slot="footer" class="dialog-footer">
            <el-button @click="createProblemDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="createProblem">Create</el-button>
          </div>
        </el-dialog>
        <el-dialog :visible="editProblemDialogVisible" @update:visible="val => editProblemDialogVisible = val" title="Edit Problem">
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
          <div slot="footer" class="dialog-footer">
            <el-button @click="editProblemDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="updateProblem">Update</el-button>
          </div>
      </el-tab-pane>
      <el-tab-pane label="Submissions" name="submissions">
        <el-table :data="submissions" style="width: 100%">
          <el-table-column prop="id" label="ID"></el-table-column>
          <el-table-column prop="problem_title" label="Problem Title"></el-table-column>
          <el-table-column prop="user_name" label="User Name"></el-table-column>
          <el-table-column prop="status" label="Status"></el-table-column>
          <el-table-column prop="created_at" label="Created At"></el-table-column>
          <el-table-column label="Actions">
            <template #default="{ row }">
              <el-button type="text" @click="viewSubmission(row)">View</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";

export default {
  data() {
    return {
      activeTab: "problems",
      createProblemDialogVisible: false,
      editProblemDialogVisible: false,
      newProblem: {
        title: "",
        difficulty: "easy",
      },
      selectedProblem: {
        id: "",
        title: "",
        difficulty: "",
      },
    };
  },
  computed: {
    ...mapState(["problems", "submissions"]),
  },
  methods: {
    ...mapActions([
      "createProblem",
      "updateProblem",
      "deleteProblem",
      "fetchSubmissions",
    ]),
    editProblem(problem) {
      this.selectedProblem = { ...problem };
      this.editProblemDialogVisible = true;
    },
    viewSubmission(submission) {
      // TODO: implement view submission functionality
    },
  },
  created() {
    this.fetchSubmissions();
  },
};
</script>
