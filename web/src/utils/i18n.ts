import { createI18n } from "vue-i18n";

const messages = {
  zh: {
    message: {
      hello: "你好，世界",
      welcome: "欢迎",
    },
    nav: {
        editor: "在线代码编辑器",
        'personal-center': "个人中心",
        logout: "退出登录",
        admin: "管理",
        'login-register': '登录/注册',
        home: '首页',
        problem: '题库',
        assignment: '竞赛/作业',
        status: '提交状态',
        rank: '排名',
    }
  },
  en: {
    message: {
      hello: "hello world",
    },
  },
  ja: {
    message: {
      hello: "こんにちは、世界",
    },
  },
};

const i18n = createI18n({
  locale: localStorage.getItem("locale") || "zh",
  fallbackLocale: "zh",
  legacy: false,
  messages,
});

export default i18n;