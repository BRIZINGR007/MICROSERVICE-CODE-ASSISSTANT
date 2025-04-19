import { Injectable } from '@angular/core';
import { HttpClientService } from './http-client.service';
import { catchError, firstValueFrom, Observable } from 'rxjs';
import { HttpErrorResponse } from '@angular/common/http';
import { environment } from '../../../environments/environment';
import { CodebasePayload, ExtractCodePayload } from '../interfaces/dashboard.interfaces';

@Injectable({
  providedIn: 'root'
})
export class DashboardService {
  private readonly CodeAssistantServiceEndPoint = environment.codeassistant_service_endpoint;

  constructor(private httpClientService: HttpClientService) { }

  retrieveCodeBaseData(): Promise<CodebasePayload[]> {
    const observable = this.httpClientService.get<CodebasePayload[]>(
      `${this.CodeAssistantServiceEndPoint}/users/get-code-bases`
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
    return firstValueFrom(observable);
  }


  extractCode(payload: ExtractCodePayload): Observable<any> {
    return this.httpClientService.post(
      `${this.CodeAssistantServiceEndPoint}/code-assist/extract-code`,
      payload
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
  }
  DeleteCodeBasdeContext(codeBaseId: string): Observable<any> {
    return this.httpClientService.delete(
      `${this.CodeAssistantServiceEndPoint}/users/delete-codebase-context`,
      { codeBaseId: codeBaseId },
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
  }





}
