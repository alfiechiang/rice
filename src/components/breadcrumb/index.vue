<template>
    <div id="bread">
        <el-breadcrumb separator="/">
          <el-breadcrumb-item v-for="(item, index) in breadcrumbs" :key="item.path">
            <span>{{ item.meta.title }}</span>
          </el-breadcrumb-item>
        </el-breadcrumb>
    </div>
</template>
<script setup lang="ts">
import { ref ,watch} from "vue";
import { useRoute,type RouteLocationMatched } from "vue-router";

const breadcrumbs = ref([] as Array<RouteLocationMatched>);
const currentRoute = useRoute();


function getBreadcrumb() {
  let matched = currentRoute.matched.filter(
    (item) => item.meta && item.meta.title
  );
  
  const first = matched[0];
  breadcrumbs.value = matched.filter((item) => {
    return item.meta && item.meta.title && item.meta.breadcrumb !== false;
  });

}


getBreadcrumb();

watch(
  () => currentRoute.path,
  (path) => {
    getBreadcrumb();
  }
);

</script>
<style scope lang="scss">
#bread {
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
}
</style>