import { Component, AfterViewInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements AfterViewInit {
  title = 'go-trigger-app';

  constructor(private router: Router) {}

  ngAfterViewInit() {
    // Your existing cursor logic
    document.addEventListener('mousemove', (ev: MouseEvent) => {
      const cursor = document.getElementById('cursor');
      if (cursor) {
        cursor.style.top = `${ev.clientY}px`;
        cursor.style.left = `${ev.clientX}px`;
      }
    });
  }

  // Helper method to check if we're on the root route
  isRootRoute(): boolean {
    return this.router.url === '/';
  }
}