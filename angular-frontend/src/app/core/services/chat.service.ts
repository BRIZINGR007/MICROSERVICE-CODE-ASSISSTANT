import { Injectable } from '@angular/core';
import { HttpClientService } from './http-client.service';
import { environment } from '../../../environments/environment';
import { catchError, firstValueFrom } from 'rxjs';
import { HttpErrorResponse } from '@angular/common/http';
import { IChat } from '../interfaces/chat.service.interface';

@Injectable({
  providedIn: 'root'
})
export class ChatService {
  private readonly CodeAssistantServiceEndPoint = environment.codeassistant_service_endpoint;


  constructor(private httpClientService: HttpClientService) { }

  UnSetCodeBaseChatSignal(): void {


  }
  RetriveChatHistory(codeBaseId: string): Promise<IChat[]> {
    const observable = this.httpClientService.get<IChat[]>(
      `${this.CodeAssistantServiceEndPoint}/chat/get-all-chats`,
      { codeBaseId: codeBaseId },
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
    return firstValueFrom(observable);
  }

  ChatWithCodeBase(userQuery: string, codeBaseId: string): Promise<IChat> {
    const observable = this.httpClientService.get<IChat>(
      `${this.CodeAssistantServiceEndPoint}/code-assist/code-base-chat`,
      { query: userQuery, codeBaseId: codeBaseId },
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
    return firstValueFrom(observable);

  }
}
