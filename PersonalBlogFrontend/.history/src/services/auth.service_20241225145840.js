import apiClient from './axios'

export const authService = {
  async login(credentials) {
    const response = await apiClient.post('/auth/login', credentials)
    if (response.data.token) {
      localStorage.setItem('token', response.data.token)
    }
    return response.data
  },

  async register(userData) {
    const response = await apiClient.post('/auth/register', userData)
    return response.data
  },

  logout() {
    localStorage.removeItem('token')
  },

  async getCurrentUser() {
    const response = await apiClient.get('/auth/me')
    return response.data
  }
} 