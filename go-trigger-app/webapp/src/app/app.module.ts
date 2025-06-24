import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule, Routes } from '@angular/router';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms'; // Add this import

import { AppComponent } from './app.component';
import { NavbarComponent } from './navbar/navbar.component';
import { HeadComponent } from './head/head.component';
import { AboutComponent } from './about/about.component';
import { SkillsComponent } from './skills/skills.component';
import { ProjectsComponent } from './projects/projects.component';
import { ProjectCardComponent } from './project-card/project-card.component';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { AuthGuard } from './auth.guard'; // Add this import
import { AuthService } from './auth.service';
import { RegisterComponent } from './register/register.component'; // Add this import

const routes: Routes = [
  { path: '', component: AppComponent },
  { path: 'home', component: HomeComponent, canActivate: [AuthGuard] },
  { path: 'login', component: LoginComponent },
  { path: 'register', component: RegisterComponent },
  { path: '**', redirectTo: '' }
];
@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    HeadComponent,
    AboutComponent,
    SkillsComponent,
    ProjectsComponent,
    ProjectCardComponent,
    HomeComponent,
    LoginComponent,
    RegisterComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    FormsModule, // Add this line
    RouterModule.forRoot(routes)
  ],
  providers: [
    AuthService, // Add this service to providers
    AuthGuard // Add this guard to providers
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }