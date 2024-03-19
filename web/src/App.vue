<template>
  <router-view id="root-router" />
</template>

<style>
body {
  margin: 0px;
}

#root-router {
  min-height: calc(100vh - 151px);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>