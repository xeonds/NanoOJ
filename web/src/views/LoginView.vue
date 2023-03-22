<template>
  <div id="register">
    <h1>{{ $route.path === "/login" ? "Login" : "Register" }}</h1>
    <el-row :gutter="40">
      <el-col :span="8">
        <el-card>
          <el-form>
            <el-form-item v-if="$route.path === '/register'" label="Username">
              <el-input v-model="username"></el-input>
            </el-form-item>
            <el-form-item label="Email">
              <el-input v-model="email"></el-input>
            </el-form-item>
            <el-form-item label="Password">
              <el-input type="password" v-model="password"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button
                type="primary"
                v-if="$route.path === '/login'"
                @click.prevent="login"
                >Login</el-button
              >
              <el-button type="primary" v-else @click.prevent="register"
                >Register</el-button
              >
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
      <el-col :span="2"></el-col>
      <el-col :span="14">
        <h2>Nano OJ</h2>
        <p>
          Our system provides a platform for users to solve programming problems
          and improve their coding skills. You can submit your solutions and get
          feedback on your code's correctness and efficiency. Join our community
          and start solving problems today!
        </p>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: "",
      email: "",
      password: "",
    };
  },
  methods: {
    async login() {
      try {
        const response = await this.axios.post("/user/login", {
          email: this.email,
          password: this.password,
        });
        const data = response.data;
        if (data.token != null) {
          this.$store.status.token = data.token;
          this.$store.status.isLoggedIn = true;
          this.$router.push("/");
        }
      } catch (error) {
        console.error(error);
      }
    },
    async register() {
      try {
        const response = await this.axios.post("/user/register", {
          username: this.username,
          email: this.email,
          password: this.password,
        });
        const data = response.data;
        if (data.success) {
          this.$router.push("/login");
        }
      } catch (error) {
        console.error(error);
      }
    },
  },
};
</script>

<style scoped>
#register {
  height: 50vh;
}

.el-card {
  width: 24rem;
}
</style>
