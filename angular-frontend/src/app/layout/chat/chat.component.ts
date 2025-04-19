import { Component, effect } from '@angular/core';
import { HeaderComponent } from '../header/header.component';
import { ChatService } from '../../core/services/chat.service';
import { ChatSkeletonComponent } from '../../shared/components/chat-skeleton/chat-skeleton.component';
import { ChatMessageComponent } from '../chat-message/chat-message.component';
import { ActivatedRoute, Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { RefrenceTraversalComponent } from '../refrence-traversal/refrence-traversal.component';
import { IChat } from '../../core/interfaces/chat.service.interface';

@Component({
  selector: 'app-chat',
  imports: [HeaderComponent, ChatSkeletonComponent, ChatMessageComponent, FormsModule, CommonModule, RefrenceTraversalComponent],
  templateUrl: './chat.component.html',
  styleUrl: './chat.component.scss'
})
export class ChatComponent {
  isLoading: boolean = true;
  codeChats: IChat[] = [];
  userQuestion: string = '';
  codeBaseId: string = "";
  codeBaseName: string = "";
  protected showReferencesTraversal: boolean = false;
  constructor(
    private chatService: ChatService,
    private route: ActivatedRoute
  ) { }

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.codeBaseId = params['codeBaseId'];
      this.codeBaseName = params['codeBaseName'];
      this.ComponentDataSetup();
    });
  }


  protected async ComponentDataSetup() {
    this.codeChats = await this.chatService.RetriveChatHistory(this.codeBaseId);
    console.log(this.codeChats);
    this.isLoading = false;
    if (!this.codeChats) {
      this.populateDummyChats();
    }

  }
  SetshowRefernceTraversal(value: boolean) {
    this.showReferencesTraversal = value;
  }

  appendUserQuestionToChatHistory(question: string) {
    const newChat: IChat = {
      chat_id: "123",
      user_id: "123",
      code_base_id: this.codeBaseId,
      code_base_name: this.codeBaseName,
      ai_answer: "",
      user_question: question,
      references: []
    };
    this.codeChats.push(newChat);
    console.log(this.codeChats);
  }
  updateChat(chat: IChat) {
    if (this.codeChats.length > 0) {
      this.codeChats.pop();
    }
    this.codeChats.push(chat);
  }

  protected async handleSend() {
    if (this.userQuestion.trim()) {
      this.appendUserQuestionToChatHistory(this.userQuestion)
      const chat = await this.chatService.ChatWithCodeBase(this.userQuestion, this.codeBaseId);
      this.updateChat(chat);
    }
  }
  protected populateDummyChats() {

    this.codeChats = [{
      chat_id: "123",
      user_id: "123",
      code_base_id: this.codeBaseId,
      code_base_name: this.codeBaseName,
      ai_answer: "",
      user_question: "Welcome To Briznigr's  Code Assistant . Please Chat On .",
      references: []
    }]
  }

}
