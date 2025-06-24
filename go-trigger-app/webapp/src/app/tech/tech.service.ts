import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Technology } from './tech.model';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class TechService {
  private apiUrl = environment.apiUrl;

  constructor(private readonly httpClient: HttpClient) { }

  getTechnologies(): Observable<Technology[]> {
    return this.httpClient.get<Technology[]>(`${this.apiUrl}/api/technologies`);
  }
}