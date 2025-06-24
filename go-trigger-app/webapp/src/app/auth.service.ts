import { Injectable } from '@angular/core';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private isAuthenticated = false;

  constructor(private router: Router) {}

  login(username: string, password: string): boolean {
    // Replace with actual authentication logic
    if (username === 'demo' && password === 'demo') { // Temporary demo check
      this.isAuthenticated = true;
      return true;
    }
    return false;
  }

  logout(): void {
    this.isAuthenticated = false;
    this.router.navigate(['/']);
  }

  isLoggedIn(): boolean {
    return this.isAuthenticated;
  }

register(username: string, email: string, password: string): boolean {
  // Replace with actual registration logic
  // For now, we just simulate registration with localStorage
  let users: Array<{ username: string; email: string; password: string }> = [];
  const usersJson = localStorage.getItem('users');
  if (usersJson) {
    users = JSON.parse(usersJson);
  }

  if (users.some((user) => user.username === username)) {
    return false;
  }

  users.push({ username, email, password });
  localStorage.setItem('users', JSON.stringify(users));
  this.isAuthenticated = true;
  return true;
}
  getCurrentUser(): string | null {
    // Replace with actual user retrieval logic
    if (this.isAuthenticated) {
      return 'demo'; // Temporary demo user
    }
    return null;
  }
}