
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

import { Component, HostListener } from '@angular/core'; 

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent {
  isDropdownOpen = false;

  constructor(private router: Router) {}

  toggleDropdown(event: Event): void {
    event.preventDefault();
    event.stopPropagation();
    this.isDropdownOpen = !this.isDropdownOpen;
  }

  @HostListener('document:click', ['$event'])
  closeDropdown(event: Event): void {
    if (!(event.target as HTMLElement).closest('.navbar-dropdown')) {
      this.isDropdownOpen = false;
    }
  }

  @HostListener('document:keydown.escape')
  onEscapeKey(): void {
    this.isDropdownOpen = false;
  }

  isActive(route: string): boolean {
    return this.router.url === route;
  }
}