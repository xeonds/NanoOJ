<template>
  <div>
    <h1>{{ problemTitle }}</h1>
    <p>{{ problemNote }}</p>
    <el-form>
      <el-form-item>
        <vue3-ace-editor v-model="code" mode="javascript" theme="chrome" width="100%" height="500px"></vue3-ace-editor>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitCode">Commit</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import axios from 'axios'

export default ({
  name: 'ProblemView',
  components: {
  },
  computed: {
    problemData() {
      return this.$store.state.problemData
    },
    problemTitle() {
      const id = this.$route.params.id
      return this.problemData[id].title
    },
    problemNote() {
      const id = this.$route.params.id
      return this.problemData[id].note
    }
  },
  setup() {
    const code = ''
    const submitCode = () => {
      axios.post('/api/v1/judge', { code: code.value })
        .then(response => {
          console.log(response.data)
        })
        .catch(error => {
          console.log(error)
        })
    }
    return {
      code,
      submitCode
    }
  }
})
</script>

<style scoped>
</style>
