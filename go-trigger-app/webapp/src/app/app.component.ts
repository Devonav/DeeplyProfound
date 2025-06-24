import { Component, AfterViewInit } from '@angular/core';
<<<<<<< HEAD
import { Router } from '@angular/router';
=======
>>>>>>> 294e5ea1203622e3f59a55780ea370922c4dcad8

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements AfterViewInit {
  title = 'go-trigger-app';

<<<<<<< HEAD
  constructor(private router: Router) {}

  ngAfterViewInit() {
    // Your existing cursor logic
=======
  ngAfterViewInit() {
>>>>>>> 294e5ea1203622e3f59a55780ea370922c4dcad8
    document.addEventListener('mousemove', (ev: MouseEvent) => {
      const cursor = document.getElementById('cursor');
      if (cursor) {
        cursor.style.top = `${ev.clientY}px`;
        cursor.style.left = `${ev.clientX}px`;
      }
    });
  }
<<<<<<< HEAD

  // Helper method to check if we're on the root route
  isRootRoute(): boolean {
    return this.router.url === '/';
  }
}
=======
}
>>>>>>> 294e5ea1203622e3f59a55780ea370922c4dcad8
