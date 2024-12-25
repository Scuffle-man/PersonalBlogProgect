import { defineStore } from 'pinia'
import { authService } from '@/services/auth.service'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token'),
    user: JSON.parse(localStorage.getItem('user') || 'null'),
  }),
  
  getters: {
    isAuthenticated: (state) => !!state.token,
  },
  
  actions: {
    async login(credentials) {
      const response = await authService.login(credentials);
      this.token = response.token;
      this.user = response.user;
      localStorage.setItem('user', JSON.stringify(response.user));
      return response;
    },
    
    logout() {
      authService.logout();
      this.token = null;
      this.user = null;
      localStorage.removeItem('user');
    }
  }
}) 