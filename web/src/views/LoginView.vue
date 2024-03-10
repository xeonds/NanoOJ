<template>
  <div id="register">
    <h1>{{ route.path === "/login" ? "Login" : "Register" }}</h1>
    <el-row :gutter="40">
      <el-col :span="8">
        <el-card>
          <el-form>
            <el-form-item v-if="route.path === '/register'" label="Username">
              <el-input v-model="username"></el-input>
            </el-form-item>
            <el-form-item label="Email">
              <el-input v-model="email"></el-input>
            </el-form-item>
            <el-form-item label="Password">
              <el-input type="password" v-model="password"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" v-if="route.path === '/login'" @click.prevent="login">Login</el-button>
              <el-button type="primary" v-else @click.prevent="register">Register</el-button>
            </el-form-item>
          </el-form>
          <p v-if="route.path === '/login'">Have no account? Head to <el-link type="primary"
              @click="router.push('/register')">register</el-link>.</p>
          <p v-else>Already have an account? Head to <el-link type="primary"
              @click="router.push('/login')">login</el-link>.</p>
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

<script lang="ts" setup>
import api from '@/api';
import { setToken } from '@/utils/login';

const router = useRouter();
const route = useRoute();

const username = ref("");
const email = ref("");
const password = ref("");

const login = async () => {
  try {
    const response = await api.login({ email: email.value, password: password.value });
    setToken(response.data.value.token);
    router.push("/");
  } catch (error) {
    console.log(error);
  }
};

const register = async () => {
  try {
    await api.register({ username: username.value, email: email.value, password: password.value });
    router.push("/login");
  } catch (error) {
    console.error(error);
  }
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
