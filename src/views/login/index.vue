<template>
   <div id="login">
      <el-form class="login-form">
         <h1 style="text-align: center;margin-bottom: 20px;">AXC 系統</h1>
         <el-form-item>
            <div class="p-2">
               <el-icon>
                  <User />
               </el-icon>
            </div>
            <el-input class="flex-1" placeholder="用戶名" />
         </el-form-item>
         <el-form-item>
            <div class="p-2">
               <el-icon><Lock /></el-icon>
            </div>
            <el-input class="flex-1" placeholder="密碼" />
         </el-form-item>
         <el-button type="primary" class="w-full" >登入</el-button>
      </el-form>
   </div>
</template>
<script setup lang="ts">

</script>
<style scoped lang="scss">
#login {
   width: 100%;
   height: 100%;
   background-color: #2D3A4B;
   color: #fff;
   display: flex;
   justify-content: center;

   .login-form {
      width: 500px;
      height: 200px;
      margin-top: 200px;
      
      .el-form-item {
         background: rgb(0 0 0 / 10%);
         border: 1px solid rgb(255 255 255 / 10%);
         border-radius: 5px;
      }

      .el-input {
         background: transparent;
         // 子组件 scoped 无效，使用 :deep
         :deep(.el-input__wrapper) {
            padding: 0;
            background: transparent;
            box-shadow: none;

            .el-input__inner {
               color: #fff;
               background: transparent;
               border: 0;
               border-radius: 0;
               caret-color: #fff;

               &:-webkit-autofill {
                  box-shadow: 0 0 0 1000px transparent inset !important;
                  -webkit-text-fill-color: #fff !important;
               }

               // 设置输入框自动填充的延迟属性
               &:-webkit-autofill,
               &:-webkit-autofill:hover,
               &:-webkit-autofill:focus,
               &:-webkit-autofill:active {
                  transition: color 99999s ease-out, background-color 99999s ease-out;
                  transition-delay: 99999s;
               }
            }
         }
      }


   }

}





.w-full{
   width: 100%;
}

.p-2 {
   padding: 0.5rem;
}

.flex-1 {
   flex: 1 1 0%;
}
</style>