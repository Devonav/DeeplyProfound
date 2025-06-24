import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  username = '';
  email = '';
  password = '';
  confirmPassword = '';
  error = '';

  constructor(
    private authService: AuthService,
    private router: Router
  ) {}

  onSubmit(): void {
    if (this.password !== this.confirmPassword) {
      this.error = 'Passwords do not match';
      return;
    }

    this.authService.register(this.username, this.password).subscribe({
      next: (response) => {
        if (response && response.id) {
          // On successful registration, redirect to the login page
          this.router.navigate(['/login']);
        } else {
          this.error = 'Registration failed. Please try again.';
        }
      },
      error: (err) => {
        this.error = 'Registration failed. Please try again.';
      }
    });
  }
}