

<template>
  <el-config-provider :locale="locale">
    <RouterView />
  </el-config-provider>
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router'
import zhTw from 'element-plus/dist/locale/zh-tw.mjs'
const locale = zhTw 
</script>

<style scoped></style>
