<template>
  <div id="register">
    <h1 style="padding-inline: 2rem">{{ route.path === "/login" ? t('message.login') : t('message.register') }}</h1>
    <div style="padding: 4rem; display: flex;flex-flow: column; justify-content: center;">
      <el-row :gutter="40">
        <el-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8">
          <el-card>
            <el-form>
              <el-form-item v-if="route.path === '/register'" :label="t('user.name')">
                <el-input v-model="username"></el-input>
              </el-form-item>
              <el-form-item :label="t('user.email')">
                <el-input v-model="email"></el-input>
              </el-form-item>
              <el-form-item :label="t('message.password')">
                <el-input type="password" v-model="password"></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" v-if="route.path === '/login'" @click.prevent="login">{{ t('message.login') }}</el-button>
                <el-button type="primary" v-else @click.prevent="register">{{ t('message.register') }}</el-button>
              </el-form-item>
            </el-form>
            <p v-if="route.path === '/login'">{{ t('message.have-no-account') }}<el-link type="primary"
                @click="router.push('/register')">{{ t('message.register') }}</el-link></p>
            <p v-else>{{ t('message.already-have-account') }}<el-link type="primary"
                @click="router.push('/login')">{{ t('message.login') }}</el-link></p>
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
import { handleHttp } from '@/utils/http';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const router = useRouter();
const route = useRoute();

const username = ref("");
const email = ref("");
const password = ref("");

const login = async () =>
  handleHttp(
    await api.login({ email: email.value, password: password.value } as User),
    (data: { value: { token: string; }; }) => {
      setToken(data.value.token);
      ElMessage({ message: "Login successful", type: "success" });
      router.push("/");
    },
    (err: any) => ElMessage.error("Invalid email or password: ", err)
  );

const register = async () =>
  handleHttp(
    await api.register({ username: username.value, email: email.value, password: password.value } as User),
    () => {
      router.push("/login");
    },
    (err: any) => ElMessage.error("Invalid email or password: ", err)
  );

</script>

<style scoped>
#register {
  height: calc(100vh - 10rem);
}
</style>
