import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  username = '';
  password = '';
  error = '';

  constructor(
    private authService: AuthService,
    private router: Router
  ) {}

  onSubmit(): void {
    this.authService.login(this.username, this.password).subscribe({
      next: (response) => {
        if (response && response.token) {
          this.router.navigate(['/leaks']);
        } else {
          this.error = 'Invalid username or password';
        }
      },
      error: (err) => {
        this.error = 'Invalid username or password';
      }
    });
  }
}