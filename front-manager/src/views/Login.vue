<template>
  <div class="login-container">
    <div class="login-box">
      <!-- 表单元素 -->
    <el-form status-icon ref="ruleForm" :rules="rules" :model="loginForm" label-width="100px" class="login-form">
      <el-form-item>
        <h2>登录</h2>
      </el-form-item>
      <el-form-item label="用户名" prop="username">
        <el-input type="text" placeholder="请输入用户名" v-model="loginForm.username"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input type="password" placeholder="请输入密码" v-model="loginForm.password"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" style="width: 100%; margin-bottom: 30px" @click.native.prevent="handleLogin">登录</el-button>
      </el-form-item>
    </el-form>
    </div>
  </div>
</template>

<script>
export default {
name:'Login',
data() {
    return {
      loginForm: {
        username: 'admin',
        password: 'admin'
      },

      rules: {
        username: [{ required: true, message: '请输入登录名称', trigger: 'blur' },{ min:5,max:10, message:'长度在5到10个字符', trigger: 'blur' }],
        password: [{ required: true, message: '请输入密码', trigger: 'blur' },{ min:5,max:10, message:'长度在5到10个字符', trigger: 'blur' }]
      }
    }
  },
methods:{
  handleLogin(){
    this.$refs.ruleForm.validate(async (valid)=>{
      if(valid){
       let {data:res} =  await this.$API.user.login(this.loginForm)
       if(res.code==0){
        this.$message.success('登录成功')
        // 存储token
       localStorage.setItem('token', res.data.token)
        // 路由跳转到home
        this.$router.push('/home')
       }else{
        this.$message.error('登录失败')
       }
      }
    })
  }
}
}
</script>

<style scoped lang="less">
.login-container {
  height: 100%;
  background-color: #2b4b6b;
  .login-box{
    width: 600px;
    height: 300px;
    background-color: #eee;
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%,-50%);
    border-radius: 10px;
    .login-form {
    width: 520px;
    max-width: 100%;
  }
  .el-button {
    margin-left: 0;
  }
  }
  
}
</style>
