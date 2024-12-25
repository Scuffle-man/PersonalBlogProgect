<template>
  <div class="login-container">
    <div class="login-card">
      <h1 class="welcome-text">WELCOME</h1>
      
      <form @submit.prevent="handleLogin" class="login-form">
        <div class="form-group">
          <input 
            type="email" 
            v-model="email"
            placeholder="Your email"
            class="form-input"
          />
        </div>

        <div class="form-group">
          <input 
            type="password"
            v-model="password"
            placeholder="Your password"
            class="form-input"
          />
        </div>

        <div class="form-options">
          <label class="remember-me">
            <input type="checkbox" v-model="rememberMe">
            <span>Remember me</span>
          </label>
          <a href="#" class="forgot-password">Forgot password?</a>
        </div>

        <button type="submit" class="login-button">
          Login
        </button>

        <p class="register-link">
          Don't have an account? 
          <router-link to="/register">Register</router-link>
        </p>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const email = ref('')
const password = ref('')
const rememberMe = ref(false)
const router = useRouter()

const handleLogin = async () => {
  // 实现登录逻辑
}
</script>

<style lang="scss" scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #1a2a6c, #b21f1f, #fdbb2d);
  position: relative;
  overflow: hidden;

  &::before {
    content: '';
    position: absolute;
    width: 100%;
    height: 100%;
    background-image: url('@/assets/images/cityscape.jpg');
    background-size: cover;
    background-position: center;
    opacity: 0.8;
  }
}

.login-card {
  width: 100%;
  max-width: 400px;
  padding: 2rem;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-radius: 1rem;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  animation: fadeIn 0.5s ease-out;
}

.welcome-text {
  font-size: 2rem;
  color: white;
  text-align: center;
  margin-bottom: 2rem;
  font-weight: 600;
  letter-spacing: 2px;
}

.form-input {
  width: 100%;
  padding: 1rem;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 2rem;
  color: white;
  margin-bottom: 1rem;
  transition: all 0.3s ease;

  &:focus {
    outline: none;
    border-color: rgba(255, 255, 255, 0.5);
    background: rgba(255, 255, 255, 0.2);
  }
}

.form-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 1rem 0;
  color: white;
}

.login-button {
  width: 100%;
  padding: 1rem;
  background: white;
  border: none;
  border-radius: 2rem;
  color: #1a2a6c;
  font-weight: 600;
  cursor: pointer;
  transition: transform 0.3s ease;

  &:hover {
    transform: translateY(-2px);
  }
}

.register-link {
  text-align: center;
  margin-top: 1rem;
  color: white;

  a {
    color: white;
    text-decoration: underline;
  }
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style> 