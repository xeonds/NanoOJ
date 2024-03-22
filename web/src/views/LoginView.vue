<template>
  <div id="register">
    <h1 style="padding-inline: 2rem">{{ route.path === "/login" ? "Login" : "Register" }}</h1>
    <div style="padding: 4rem; display: flex;flex-flow: column; justify-content: center;">
      <el-row :gutter="40">
        <el-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8">
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
        <el-col :xs="0" :sm="0" :md="2" :lg="2" :xl="2"></el-col>
        <el-col :xs="24" :sm="24" :md="14" :lg="14" :xl="14">
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
  </div>
  <FooterBox id="footer" :msg="'X-OJ'" />
</template>

<script lang="ts" setup>
import api from '@/api';
import FooterBox from '@/components/FooterBox.vue';
import { setToken } from '@/utils/login';
import { User } from '@/model';

const router = useRouter();
const route = useRoute();

const username = ref("");
const email = ref("");
const password = ref("");

const login = async () => {
  try {
    const response = await api.login({ email: email.value, password: password.value } as User);
    setToken(response.data.value.token);
    ElMessage({ message: "Login successful", type: "success" })
    router.push("/");
  } catch (error) {
    ElMessage.error("Invalid email or password");
  }
};

const register = async () => {
  try {
    await api.register({ username: username.value, email: email.value, password: password.value } as User);
    router.push("/login");
  } catch (error) {
    ElMessage.error("Invalid email or password");
  }
};
</script>

<style scoped>
#register {
  height: calc(100vh - 10rem);
}
</style>
