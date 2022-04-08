<template>
  <div>
    <nav>
      <router-link to="/">主页</router-link> |
      <router-link to="/problemset">问题</router-link> |
      <router-link to="/category">来源/分类</router-link> |
      <router-link to="/contest">竞赛/作业</router-link> |
      <router-link to="/status">状态</router-link> |
      <router-link to="/ranklist">排名</router-link> |
      <router-link to="/about">常见问答</router-link> |
      <router-link to="/user">Username</router-link>
    </nav>
    <router-view />
  </div>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}
</style>
