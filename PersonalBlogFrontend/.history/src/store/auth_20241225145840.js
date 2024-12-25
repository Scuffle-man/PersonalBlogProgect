import { defineStore } from 'pinia'
import { authService } from '@/services/auth.service'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    token: localStorage.getItem('token')
  }),
  
  getters: {
    isAuthenticated: state => !!state.token,
    userProfile: state => state.user
  },
  
  actions: {
    async login(credentials) {
      const data = await authService.login(credentials)
      this.token = data.token
      this.user = data.user
    },
    
    async logout() {
      authService.logout()
      this.token = null
      this.user = null
    },
    
    async fetchCurrentUser() {
      try {
        const user = await authService.getCurrentUser()
        this.user = user
      } catch (error) {
        this.logout()
      }
    }
  }
}) 